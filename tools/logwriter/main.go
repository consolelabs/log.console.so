// logwriter drafts a Console Labs monthly lab-journal post from real repo signal,
// in the lab voice (docs/voice.md). Hybrid: it gathers the month's merged-PR
// signal AND drafts the narration via the Anthropic API. Draft-only: it writes a
// vault-ready Obsidian markdown file; it never publishes.
//
// Usage:
//
//	logwriter 2026-05                          # draft May 2026 to ./logwriter-out/
//	logwriter 2026-05 --dry-run                # gather signal only, emit stub (no LLM)
//	logwriter 2026-05 --out vault/blog         # write into the content vault
//	logwriter 2026-05 --owner consolelabs --repo tieubao/neko
//
// The Anthropic key is resolved at run time from 1Password
// (op://Toolkit/anthropic-api-key/credential) or the ANTHROPIC_API_KEY env var.
// It is never hardcoded.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

const (
	defaultModel  = "claude-sonnet-4-6"
	opKeyRef      = "op://Toolkit/anthropic-api-key/credential"
	voiceRelPath  = "docs/voice.md" // resolved relative to repo root (cwd or --repo-root)
	anthropicURL  = "https://api.anthropic.com/v1/messages"
	anthropicVers = "2023-06-01"
)

type pr struct {
	Title      string `json:"title"`
	URL        string `json:"url"`
	ClosedAt   string `json:"closedAt"`
	Repository struct {
		Name          string `json:"name"`
		NameWithOwner string `json:"nameWithOwner"`
	} `json:"repository"`
}

type stringList []string

func (s *stringList) String() string { return strings.Join(*s, ",") }
func (s *stringList) Set(v string) error {
	*s = append(*s, v)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "logwriter:", err)
		os.Exit(1)
	}
}

func run() error {
	var owners, repos stringList
	out := flag.String("out", "logwriter-out", "output directory for the draft")
	model := flag.String("model", defaultModel, "Anthropic model for drafting")
	backend := flag.String("backend", "api", "LLM backend: api (Anthropic API via op key) or claude (claude CLI)")
	repoRoot := flag.String("repo-root", ".", "repo root holding docs/voice.md")
	dryRun := flag.Bool("dry-run", false, "gather signal + emit a stub; skip the LLM draft")
	flag.Var(&owners, "owner", "GitHub org to scan (repeatable; default consolelabs)")
	flag.Var(&repos, "repo", "extra owner/repo to scan (repeatable; default tieubao/neko, tieubao/console-labs)")

	// The month is the first positional; flags may follow it. Go's flag package
	// stops at the first non-flag arg, so pull the month out explicitly.
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		return fmt.Errorf("usage: logwriter <YYYY-MM> [flags]")
	}
	monthArg := os.Args[1]
	flag.CommandLine.Parse(os.Args[2:])

	month, err := time.Parse("2006-01", monthArg)
	if err != nil {
		return fmt.Errorf("bad month %q (want YYYY-MM): %w", monthArg, err)
	}
	if len(owners) == 0 {
		owners = stringList{"consolelabs"}
	}
	if len(repos) == 0 {
		repos = stringList{"tieubao/neko", "tieubao/console-labs"}
	}

	start := month
	end := month.AddDate(0, 1, 0).AddDate(0, 0, -1) // last day of month
	rng := fmt.Sprintf("%s..%s", start.Format("2006-01-02"), end.Format("2006-01-02"))

	signal, err := gather(owners, repos, rng)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "gathered %d merged PRs across %d sources for %s\n",
		len(signal), len(owners)+len(repos), month.Format("January 2006"))

	byRepo := groupByRepo(signal)
	var body string
	mode := "drafted"
	if *dryRun || len(signal) == 0 {
		mode = "stub"
		body = stub(month, byRepo)
		if len(signal) == 0 {
			fmt.Fprintln(os.Stderr, "no signal found; emitting stub (degraded mode)")
		}
	} else {
		voice, err := os.ReadFile(*repoRoot + "/" + voiceRelPath)
		if err != nil {
			return fmt.Errorf("read voice spec: %w (run from repo root or pass --repo-root)", err)
		}
		system, user := buildPrompt(string(voice), month, byRepo)
		switch *backend {
		case "claude":
			body, err = draftClaude(system, user)
		default:
			body, err = draftAPI(*model, system, user)
		}
		if err != nil {
			return fmt.Errorf("draft via %s backend: %w (use --dry-run for stub mode)", *backend, err)
		}
		body = sanitizeBody(body)
	}

	post := frontMatter(month) + "\n" + body + "\n"
	if err := os.MkdirAll(*out, 0o755); err != nil {
		return err
	}
	path := fmt.Sprintf("%s/console-log-%s.md", *out, month.Format("2006-01"))
	if err := os.WriteFile(path, []byte(post), 0o644); err != nil {
		return err
	}
	fmt.Printf("%s draft written: %s\n", mode, path)
	return nil
}

// gather runs gh search prs per source and merges the results.
func gather(owners, repos stringList, rng string) ([]pr, error) {
	seen := map[string]bool{}
	var all []pr
	add := func(args ...string) error {
		base := []string{"search", "prs", "--merged-at", rng, "--limit", "200",
			"--json", "title,url,repository,closedAt"}
		out, err := exec.Command("gh", append(args, base...)...).Output()
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				return fmt.Errorf("gh %v: %s", args, strings.TrimSpace(string(ee.Stderr)))
			}
			return fmt.Errorf("gh %v: %w", args, err)
		}
		var got []pr
		if err := json.Unmarshal(out, &got); err != nil {
			return fmt.Errorf("parse gh output: %w", err)
		}
		for _, p := range got {
			if !seen[p.URL] {
				seen[p.URL] = true
				all = append(all, p)
			}
		}
		return nil
	}
	for _, o := range owners {
		if err := add("--owner", o); err != nil {
			return nil, err
		}
	}
	for _, r := range repos {
		if err := add("--repo", r); err != nil {
			return nil, err
		}
	}
	sort.Slice(all, func(i, j int) bool { return all[i].ClosedAt < all[j].ClosedAt })
	return all, nil
}

func groupByRepo(prs []pr) map[string][]pr {
	m := map[string][]pr{}
	for _, p := range prs {
		key := p.Repository.NameWithOwner
		if key == "" {
			key = p.Repository.Name
		}
		m[key] = append(m[key], p)
	}
	return m
}

func signalText(byRepo map[string][]pr) string {
	repos := make([]string, 0, len(byRepo))
	for r := range byRepo {
		repos = append(repos, r)
	}
	sort.Strings(repos)
	var b strings.Builder
	for _, r := range repos {
		fmt.Fprintf(&b, "\n## %s (%d merged PRs)\n", r, len(byRepo[r]))
		for _, p := range byRepo[r] {
			fmt.Fprintf(&b, "- %s (%s)\n", p.Title, p.URL)
		}
	}
	return b.String()
}

// sanitizeBody strips wrapper cruft an LLM may add despite instructions: a
// code fence wrapping the whole reply (which also drops any trailing meta-note
// outside it) and a leading YAML front-matter block (the tool owns front matter).
func sanitizeBody(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```") {
		rest := s[3:]
		if nl := strings.IndexByte(rest, '\n'); nl >= 0 {
			rest = rest[nl+1:] // drop the ```lang tag line
		}
		if j := strings.Index(rest, "```"); j >= 0 {
			s = strings.TrimSpace(rest[:j])
		}
	}
	if strings.HasPrefix(s, "---") {
		if end := strings.Index(s[3:], "\n---"); end >= 0 {
			s = strings.TrimSpace(s[3+end+4:])
		}
	}
	return s
}

func frontMatter(month time.Time) string {
	return fmt.Sprintf(`---
tags: [consolelabs, work, monthly]
title: "Console Log: %s"
date: %s
description:
authors: [neko-san]
toc: false
type:
---`, month.Format("January 2006"), month.Format("2006-01-02"))
}

// stub is the degraded / --dry-run output: the assembled signal + section stubs.
func stub(month time.Time, byRepo map[string][]pr) string {
	return fmt.Sprintf(`<!-- STUB: assembled signal only, no narration. Run without --dry-run to draft. -->

## TL;DR

_(narrate the month here)_

## Shipped
%s

## Tinkering

## Thinking

## Next
`, signalText(byRepo))
}

// buildPrompt assembles the system (voice + rules) and user (task + signal) prompts.
func buildPrompt(voice string, month time.Time, byRepo map[string][]pr) (system, user string) {
	system = voice + "\n\nYou draft the Console Labs monthly lab journal in EXACTLY this voice. " +
		"Start your reply directly with the post content (a lead paragraph or the first '## ' section). " +
		"Do NOT add YAML front matter and do NOT wrap the whole reply in a code fence. " +
		"Use the section skeleton from the voice doc (TL;DR, Shipped, Tinkering, Thinking, Next). " +
		"Drop any section with nothing real to say; never pad. Narrate ONLY from the signal given; " +
		"do not invent shipped features. Keep it skimmable and in-voice. No em dashes."
	user = fmt.Sprintf("Draft the Console Labs monthly log for %s.\n\nThe month's real signal (merged PRs grouped by repo):\n%s\n\nFor the Shipped section, group the notable work and add a one-line 'why it matters' where you can infer it from the titles. If a repo has only chore/dependency PRs, summarize rather than list.",
		month.Format("January 2006"), signalText(byRepo))
	return system, user
}

// draftAPI drafts via the Anthropic Messages API (key from op, in-process HTTP).
func draftAPI(model, system, user string) (string, error) {
	key, err := apiKey()
	if err != nil {
		return "", err
	}
	reqBody, _ := json.Marshal(map[string]any{
		"model":      model,
		"max_tokens": 4000,
		"system":     system,
		"messages":   []map[string]string{{"role": "user", "content": user}},
	})
	// In-process HTTP keeps the key in a header, never on any process argv.
	req, err := http.NewRequest("POST", anthropicURL, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("x-api-key", key)
	req.Header.Set("anthropic-version", anthropicVers)
	req.Header.Set("content-type", "application/json")
	httpResp, err := (&http.Client{Timeout: 120 * time.Second}).Do(req)
	if err != nil {
		return "", fmt.Errorf("anthropic call: %w", err)
	}
	defer httpResp.Body.Close()
	out, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}
	var resp struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(out, &resp); err != nil {
		return "", fmt.Errorf("parse anthropic response: %w", err)
	}
	if resp.Error != nil {
		return "", fmt.Errorf("anthropic: %s", resp.Error.Message)
	}
	if len(resp.Content) == 0 {
		return "", fmt.Errorf("anthropic: empty response")
	}
	return strings.TrimSpace(resp.Content[0].Text), nil
}

// draftClaude drafts via the claude CLI (headless). Works where the Anthropic
// API egress is blocked. Strips a trailing global-config canary line if present.
func draftClaude(system, user string) (string, error) {
	cmd := exec.Command("claude", "-p")
	cmd.Stdin = strings.NewReader(system + "\n\n" + user)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("claude cli: %s", strings.TrimSpace(string(ee.Stderr)))
		}
		return "", err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for len(lines) > 0 {
		last := strings.TrimSpace(lines[len(lines)-1])
		if last == "" || strings.HasPrefix(last, "🐱") {
			lines = lines[:len(lines)-1]
			continue
		}
		break
	}
	return strings.TrimSpace(strings.Join(lines, "\n")), nil
}

// apiKey resolves the Anthropic key from env, else from 1Password. Never hardcoded.
func apiKey() (string, error) {
	if k := os.Getenv("ANTHROPIC_API_KEY"); k != "" {
		return k, nil
	}
	out, err := exec.Command("op", "read", opKeyRef).Output()
	if err != nil {
		return "", fmt.Errorf("no ANTHROPIC_API_KEY and op read failed: %w", err)
	}
	k := strings.TrimSpace(string(out))
	if k == "" {
		return "", fmt.Errorf("resolved an empty Anthropic key")
	}
	return k, nil
}
