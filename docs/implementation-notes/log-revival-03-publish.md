# Implementation notes, sub-goal 03 (publish flow)

Deltas from `_meta/megagoals/log-revival/goals/03-publish-flow.md`. The spec held the
design; this records what it did not pin down.

## 2026-06-25 standalone post file, not a section in the rolling log.md

- Context: the spec said "stage a drafted post into the content vault". It did not
  say WHERE, and the vault has two conflicting shapes: the historical monthly reports
  are `## Month Year` SECTIONS inside one rolling `vault/blog/log.md`, but
  `docs/voice.md`'s skeleton + the `New Post` templater describe a STANDALONE post
  (front matter `title: "Console Log: <Month>"`, `type:`).
- Decision: stage a standalone file at `blog/console-log-YYYY-MM.md`.
- Why: it matches the voice.md skeleton + templater, and it is far more robust than
  string-inserting a new `##` section into a rolling file (no fragile anchor matching,
  clean diff, trivial revert). One post = one file = one reviewable PR.
- Impact: a future content decision may want these listed/linked from the old
  `log.md` index or surfaced in the blog menu (the sample post has no `menu:` field,
  so it renders at `/blog/console-log-2026-06/` but is not auto-added to the nav).
  That is a content/site-structure call for 05/content, not the publish mechanism.

## 2026-06-25 bash, not Go

- Context: logwriter (02) is Go; Han's default for background tools is Go/Rust.
- Decision: `publish` is a POSIX-ish bash script, co-located at `tools/logwriter/publish`.
- Why: it is pure git + `gh` orchestration (clone, copy, branch, commit, push, PR),
  human-run and one-shot, not a long-running background daemon. The ponytail rung says
  the shell already does this; a Go binary would be ceremony with no perf or
  correctness gain. The Go-default is for unattended background tooling; this is an
  interactive CLI step a person runs once a month.

## 2026-06-25 mechanism proof opens a REAL content PR (then stops)

- The spec asks to "drive it end to end for a test/sample post ... open the content
  PR, STOP." The proof run opened a real PR in `consolelabs/content` (#4) using the
  privacy-trimmed sample post from 02 (non-embargoed). It is left OPEN: merging it is
  the 05 UAT step, so the mechanism test doubles as the seed for the first real post.
- The real `logwriter` June draft narrates embargoed NFT work, so it is NOT what got
  staged; the public-safe trimmed sample is.

## 2026-06-25 absolute-path fix

- First cut did `cp "$OLDPWD/$DRAFT"` to handle relative draft paths, which broke for
  absolute paths after `cd` into the clone. Fixed by resolving `DRAFT_ABS` up front.
