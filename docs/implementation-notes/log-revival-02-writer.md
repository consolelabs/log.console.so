# Implementation notes, sub-goal 02 (logwriter)

Deltas from `_meta/megagoals/log-revival/goals/02-writing-tool.md`. The spec held
the design; this records what the spec did not pin down or what building it surfaced.

## 2026-06-25 LLM backend: two paths, not one

- Context: the spec said use the `claude` CLI OR an `op://` key, decide at execution.
- Decision: support BOTH via `--backend`. Default `api` (Anthropic Messages API,
  key from `op://Toolkit/anthropic-api-key/credential`, in-process HTTP so the
  key never lands on an argv). Added a `claude` CLI backend as the fallback.
- Why: the default `api` path is right for the operator's machine, but the
  Anthropic API egress is **blocked from this build/loop environment**
  (`dial tcp ...:443: i/o timeout`, while `gh`/`op` network fine). The `claude`
  backend works there, so it is how the proof draft was produced.
- Impact: the tool is portable across both environments. The `claude` backend
  needs a trailing-canary + meta-note strip (see `sanitizeBody`).

## 2026-06-25 gh flag: `--merged-at`, not `--merged`

- `gh search prs --merged` is a boolean (merged state); the date range needs
  `--merged-at`. The first cut used `--merged <range>` and silently gathered 0
  PRs. Fixed to `--merged-at`. Constraint the spec did not mention.

## 2026-06-25 privacy gate on the committed sample

- The real June draft narrated the unannounced NFT migration (Safe multisig,
  a specific restored token) and named specific security findings. `log.console.so`
  is a PUBLIC repo and `docs/voice.md` forbids leaking embargoed work.
- Decision: commit only a trimmed `testdata/sample-output-excerpt.md` (voice +
  structure, non-embargoed sections). The full draft stays on the operator's box.

## 2026-06-25 build: GOROOT mismatch

- `go build` failed with `compile: version "go1.24.2" does not match go tool
  version "go1.26.4"` (PATH go is brew's, GOROOT pinned to mise). Build + test
  with `env -u GOROOT`. Documented in the README + proof.

## Off-spec but in-scope

- Signal is merged PRs only; releases + issues deferred (noted in README).
- Added `sanitizeBody` + unit tests (not in the spec) because the LLM wraps
  output despite instructions; the strip is load-bearing (negative control).
