# Sub-goal 02: the monthly writing tool

**Merge policy:** auto
**Time budget:** 1 calendar day after PR-01 merges
**Proof:** a real generated draft for a recent month committed under the tool's `testdata/` or `docs/proof/`, plus a run-table of the gather + draft run (command, exit, stdout tail). The draft must read in-voice, not as a changelog dump.
**Depends on:** 01
**Branch:** feat/log-revival-02-writer
**PR base:** chore/log-revival-01-build

## Outcome

A Go CLI at `tools/logwriter/` that, given a target month, gathers the real signal from the active Console Labs repos and emits a vault-ready Obsidian monthly post drafted in the lab voice. Hybrid: it both assembles the signal AND drafts each section, and a section can be regenerated independently so the human edits rather than writes from scratch. Output is a draft file shaped for `vault/blog/`, NOT auto-published.

## Quality bar

Reads like a colleague narrating the lab's month over coffee, not a robot stapling bullet points. It pulls from real merged work, never invents shipped features, and an empty section is dropped, never padded. The voice matches `docs/voice.md` so closely a reader cannot tell a tool started it.

## How to close the loop

- Build the tool (`go build ./...` in `tools/logwriter/`), capture into a run-table at `tools/logwriter/docs/proof-of-done.md`.
- Run it for a recent month against the real signal sources and commit the generated draft as the proof artifact. Signal sources (default, refine at execution): merged PRs + releases in the target month across the active Console Labs repos (the `consolelabs` org plus `tieubao/neko`, `tieubao/console-labs`). Use `gh` for the signal.
- The draft must carry the monthly-report front matter + section skeleton from `docs/voice.md` (TL;DR, Shipped, Tinkering, Thinking, Next) and read in-voice.
- The LLM-draft step must use an available mechanism (the `claude` CLI, or an `op://` Anthropic API key resolved at run time), never a hardcoded key. Decide the mechanism at execution; if no key/CLI is reachable, the tool still emits the assembled signal + section stubs and the run-table notes the degraded mode.

**Done =** `logwriter <month>` writes a vault-ready in-voice draft from real repo signal for a recent month, the generated draft is committed as proof, and `tools/logwriter/docs/proof-of-done.md` holds the build + run table.

## Scope edges

**In:** `tools/logwriter/` (Go module, the CLI, its docs/proof), reading repo signal via `gh`, drafting against `docs/voice.md`.
**Out:** publishing / deploying (sub-goal 03), the build pipeline (01), monitoring (04).
**Not:** auto-committing to the content vault, a web UI, multi-format (weekly) output, a scheduler/daemon, hardcoding an API key, scraping non-Console-Labs repos.

## Where to look

`docs/voice.md` (the voice + skeleton), the `vault/blog/` and `vault/_templater/New Post.md` shapes (front matter + section conventions), `gh` for repo signal. New code lands in `tools/logwriter/`.

## PR body

A Go CLI (`tools/logwriter/`) that gathers a month's real signal from the active Console Labs repos and drafts a vault-ready monthly post in the lab voice (hybrid assemble + draft, section-regeneratable, draft-only). 

Verify: `go build ./...`; `logwriter <recent-month>` emits an in-voice draft from real signal; see `tools/logwriter/docs/proof-of-done.md` + the committed sample draft.

Stacked on #<PR-01>; review after it. Part of mega-goal `log-revival`.

## Notes
