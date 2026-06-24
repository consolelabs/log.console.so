# Verification: log-revival-02-writer

Behavioral change: a Go CLI (`tools/logwriter/`) that gathers a month's real
merged-PR signal and drafts a vault-ready monthly post in the lab voice.
Full detail + run-table: `tools/logwriter/docs/proof-of-done.md`.

## Green run

```
Command: env -u GOROOT go build -o logwriter .   (in tools/logwriter)
Exit:    0
Verdict: PASS

Command: env -u GOROOT go test ./...
Exit:    0   (ok: TestSanitizeBody, TestSignalGrouping)
Verdict: PASS

Command: ./logwriter 2026-06 --dry-run
Exit:    0   (gathered 155 merged PRs across 3 sources; stub written)
Verdict: PASS

Command: ./logwriter 2026-06 --backend claude --out /tmp/lw
Exit:    0   (gathered 155 PRs; drafted a full in-voice June report from real
             signal; output sanitized to single front matter, 0 fences)
Verdict: PASS  (sample excerpt: tools/logwriter/testdata/sample-output-excerpt.md)
```

## Negative control (revert -> RED -> restore)

`sanitizeBody` strips LLM wrapper cruft (code fence + stray front matter +
trailing meta-note); without it that cruft lands in the published post.

```
Step 1 GREEN:   go test ./...                 -> ok
Step 2 BREAK:   disable the fence-strip branch in sanitizeBody (`if false && ...`)
        RED:    go test ./...                 -> FAIL
                TestSanitizeBody/fence_+_front_matter_+_trailing_meta_note
Step 3 RESTORE: re-enable the branch
        GREEN:  go test ./...                 -> ok
Verdict: PASS  (the sanitization is load-bearing)
```

## Reproduce

See `tools/logwriter/docs/proof-of-done.md` "Reproduce".
