# Proof of done, logwriter (sub-goal 02)

A Go CLI that gathers a month's real merged-PR signal and drafts a vault-ready
monthly post in the Console Labs voice. Hybrid (assemble + draft), draft-only.

## Acceptance criteria

- `logwriter <month>` gathers real signal from the active Console Labs repos.
- It drafts a vault-ready, in-voice post (front matter + the voice skeleton).
- A generated sample is committed; this file holds the build + run table.

## Confirmation run-table

| # | Command | Result |
|---|---|---|
| R1 | `env -u GOROOT go build -o logwriter .` | exit 0 (GOROOT unset works around the brew/mise go mismatch) |
| R2 | `env -u GOROOT go test ./...` | `ok` (TestSanitizeBody, TestSignalGrouping) |
| R3 | `env -u GOROOT go vet .` | clean |
| R4 | `logwriter 2026-06 --dry-run` | gathered **155 merged PRs** across 3 sources; stub written with the assembled changelog grouped by repo |
| R5 | `logwriter 2026-05 --dry-run` | gathered 0 PRs -> graceful **degraded stub** (no crash) |
| R6 | `logwriter 2026-06 --backend claude` | gathered 155 PRs; drafted a full in-voice June report (TL;DR / Shipped / Tinkering / Thinking / Next), real signal, no invented features |
| R7 | output sanity | single front matter, 0 code fences, 0 trailing meta-note (after `sanitizeBody`) |

## Captured evidence

- `testdata/sample-output-excerpt.md`: a trimmed excerpt of the R6 draft. The
  full June draft narrated the unannounced NFT migration + named specific
  security findings, so per the privacy gate (`docs/voice.md`) only the
  non-embargoed sections are committed. The excerpt shows the voice + structure.

## Negative control (revert -> RED -> restore)

The wrapper-stripping in `sanitizeBody` is load-bearing (the `claude` backend
wraps output in a code fence + adds stray front matter + a trailing meta-note).

```
GREEN:   go test ./...                              -> ok
BREAK:   disable the code-fence strip in sanitizeBody (`if false && ...`)
RED:     go test ./...                              -> FAIL TestSanitizeBody/fence_+_front_matter_+_trailing_meta_note
RESTORE: re-enable the strip
GREEN:   go test ./...                              -> ok
Verdict: PASS (without the strip, raw-LLM cruft would land in the post)
```

## Reproduce

```
cd tools/logwriter
env -u GOROOT go build -o logwriter .
env -u GOROOT go test ./...
./logwriter 2026-06 --dry-run                 # signal only
./logwriter 2026-06 --backend claude --out /tmp/lw   # real draft (or default --backend api on a normal box)
```
