# Verification: log-revival-03-publish

Behavioral change: a one-command, draft-only publish flow (`tools/logwriter/publish`)
that stages a drafted monthly post into `consolelabs/content` on a branch and opens a
PR there, then stops. It must NEVER push to the live site. Companion detail:
`docs/proof/03-publish.md` + `docs/publish-runbook.md`.

## Green run

```
Command: ./publish /tmp/console-log-2026-06.md --month 2026-06 --dry-run
Exit:    0   (clones + stages; "No branch pushed, no PR opened, live site untouched")
Verdict: PASS

Command: ./publish /tmp/console-log-2026-06.md --month 2026-06
Exit:    0   (opened content PR -> https://github.com/consolelabs/content/pull/4 ; STOP)
Verdict: PASS

Command: gh pr view 4 --repo consolelabs/content --json state -q .state
Output:  OPEN   (a real draft-only PR, branch log/2026-06, post blog/console-log-2026-06.md)
Verdict: PASS
```

## Negative control (draft-only is load-bearing: nothing reaches live)

The whole safety claim is "the post never reaches live without Han's merge." If the
flow were not draft-only, the post would land on content `main` (which the deploy
watches). It does not.

```
Step GREEN (staged): the post is on the PR branch
  Command: gh api 'repos/consolelabs/content/contents/blog/console-log-2026-06.md?ref=log/2026-06' -q .name
  Output:  console-log-2026-06.md            -> present on the branch
Step CONTROL (NOT on main / NOT deployable):
  Command: gh api 'repos/consolelabs/content/contents/blog/console-log-2026-06.md?ref=main'
  Output:  404 Not Found                     -> absent from main, nothing deploys
Verdict: PASS  (the post is reachable only through the open PR; the helper cannot
         push to main or to the live gh-pages. Draft-only is enforced, not asserted.)
```

A second structural control: `grep` the `publish` script for any push to `main` /
`gh-pages` / a `--head main` merge. There is none; it only ever `git push -u origin
log/<month>` (a feature branch) and `gh pr create` (a PR, not a merge).

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so/tools/logwriter
./publish <draft.md> --month YYYY-MM --dry-run     # plan only
./publish <draft.md> --month YYYY-MM               # opens the content PR, stops
gh api 'repos/consolelabs/content/contents/blog/console-log-YYYY-MM.md?ref=main'  # -> 404
```

## Done-gate mapping

- The publish helper stages a drafted post into `consolelabs/content` and opens a
  content PR (no direct push to live): yes (content PR #4).
- The runbook documents the flow: `docs/publish-runbook.md`.
- `docs/proof/03-publish.md` carries the dry-run table + the real-run result. The live
  screenshot (after Han's merge) is deferred to sub-goal 05 (the first real post UAT),
  which merges content PR #4.
