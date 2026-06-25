# logwriter

Drafts the Console Labs monthly lab-journal post from real repo signal, in the
lab voice. Hybrid: it **gathers** the month's merged-PR signal across the active
Console Labs repos AND **drafts** the narration in the voice from
[`docs/voice.md`](../../docs/voice.md). Draft-only: it writes a vault-ready
Obsidian markdown file; it never publishes.

## Usage

```sh
# build (the GOROOT unset works around a brew/mise toolchain mismatch)
env -u GOROOT go build -o logwriter .

# draft a month (writes ./logwriter-out/console-log-YYYY-MM.md)
./logwriter 2026-06

# gather signal only, no LLM (cheap; emits a stub with the assembled changelog)
./logwriter 2026-06 --dry-run

# write straight into the content vault for editing
./logwriter 2026-06 --out ../../vault/blog

# scope the sources
./logwriter 2026-06 --owner consolelabs --repo tieubao/neko
```

Then **edit** the draft to taste and publish it through the normal flow (the
draft is a starting point, not a final post).

## How it works

1. **Gather**: `gh search prs --merged-at <month-range>` across the sources
   (default: the `consolelabs` org + `tieubao/neko` + `tieubao/console-labs`),
   deduped and grouped by repo.
2. **Draft**: builds a prompt from the voice spec + the month's signal and asks
   an LLM for the post body, then strips any wrapper cruft (code fences, stray
   front matter) and prepends the canonical front matter.
3. **Emit**: writes `console-log-YYYY-MM.md`, ready for the vault.

## LLM backends

| `--backend` | Mechanism | When |
|---|---|---|
| `api` (default) | Anthropic Messages API; key from `op://Toolkit/anthropic-api-key/credential` or `$ANTHROPIC_API_KEY`, in-process HTTP (never on argv) | normal operator runs |
| `claude` | the `claude` CLI, headless | environments where the Anthropic API egress is blocked |

The key is resolved at run time and never hardcoded.

## Flags

| Flag | Default | Meaning |
|---|---|---|
| `--dry-run` | false | gather + emit a stub; skip the LLM |
| `--backend` | `api` | `api` or `claude` |
| `--out` | `logwriter-out` | output directory |
| `--model` | `claude-sonnet-4-6` | model for the `api` backend |
| `--owner` | `consolelabs` | org to scan (repeatable) |
| `--repo` | `tieubao/neko`, `tieubao/console-labs` | extra owner/repo (repeatable) |
| `--repo-root` | `.` | where `docs/voice.md` lives |

## Scope / limits

- Signal today is **merged PRs**. Releases + issues are a future enhancement.
- It never publishes and never invents shipped features; it narrates only the
  gathered signal. An empty section is dropped, not padded.
- Output is a **draft**; a human edits before publishing. It also respects the
  privacy gate (do not publish embargoed/unannounced work).

## Publishing (the draft-only release step)

`publish` (a sibling script) takes a draft and stages it into the `consolelabs/content`
vault repo on a branch, then opens a PR there and STOPS. It never pushes to the live
site; merging the content PR is Han's hand.

```
./publish logwriter-out/console-log-2026-06.md --dry-run   # preview the plan
./publish logwriter-out/console-log-2026-06.md             # open the content PR, stop
```

Flags: `--month YYYY-MM`, `--content-repo owner/name`, `--dry-run`. Full flow:
`../../docs/publish-runbook.md`. Proof: `../../docs/proof/03-publish.md`.

Tests: `env -u GOROOT go test ./...`. Proof: `docs/proof-of-done.md`.
