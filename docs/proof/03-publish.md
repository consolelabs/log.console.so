# Proof: sub-goal 03, the reviewed (draft-only) publish flow

`tools/logwriter/publish` takes a drafted monthly post and stages it into the
`consolelabs/content` vault repo on a branch, then opens a PR there and STOPS. It
never pushes to the live site; merging the content PR + the deploy stays Han's hand.

## Dry-run (no push, no PR)

```
Command: ./publish /tmp/console-log-2026-06.md --month 2026-06 --dry-run
Exit:    0
Output:  month=2026-06 repo=consolelabs/content branch=log/2026-06 target=blog/console-log-2026-06.md
         staged diff vs main: blog/console-log-2026-06.md | 36 +++  (1 file, 36 insertions)
         DRY RUN. Would do: git commit / git push / gh pr create ...
         DRY RUN complete. No branch pushed, no PR opened, live site untouched.
Verdict: PASS  (clones + stages + shows the plan; pushes nothing)
```

## Real run (opens the content PR, then stops)

```
Command: ./publish /tmp/console-log-2026-06.md --month 2026-06
Exit:    0
Output:  staged diff vs main: blog/console-log-2026-06.md | 24 ++  (1 file, 24 insertions)
         opened content PR -> https://github.com/consolelabs/content/pull/4
         STOP. Review + merge is Han's hand; the loop does not merge this.
Verdict: PASS  (a real draft-only PR in consolelabs/content; live site untouched)
```

Content PR: **https://github.com/consolelabs/content/pull/4** (branch `log/2026-06`,
post at `blog/console-log-2026-06.md`). Left open for Han: merging it is the 05 UAT
step that publishes the first real post.

## Draft-only safety (the post never reaches live without Han)

```
Check:  after the real run, does consolelabs/content `main` contain the post?
Command: gh api repos/consolelabs/content/contents/blog/console-log-2026-06.md?ref=main
Result:  404 (the post exists ONLY on branch log/2026-06, inside the open PR)
Verdict: PASS  (nothing on main, nothing deployed; the helper cannot reach live)
```

## Live screenshot (deferred to 05)

The "post live on log.console.so" screenshot is captured at sub-goal 05 (first-post
UAT), after Han merges content PR #4 and the deploy runs. That merge is the gate
step (outward + needs the GH_PAT deploy secret); the loop prepares + opens, Han
merges. See `docs/proof/05-first-post.md` (created at 05).

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so/tools/logwriter
./logwriter 2026-06 --out logwriter-out            # draft (backend: claude or op key)
./publish logwriter-out/console-log-2026-06.md --dry-run   # preview
./publish logwriter-out/console-log-2026-06.md             # open the content PR, stop
```
