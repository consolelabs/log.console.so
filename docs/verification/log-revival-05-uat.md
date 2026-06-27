# Verification: log-revival-05-uat (the held-final, first post live)

Stateful change: publish the first monthly post and confirm the full pipeline lands it
live on log.console.so. This is the mega-goal terminus. Han authorized publish on
2026-06-28 ("update metadata of the site then proceed to publish"). Companion narrative
+ sign-off block: `docs/proof/05-first-post.md`.

## Recorded run (live publish, 2026-06-28)

```
Step 1  merge metadata PR #11 -> main                                  Exit: 0
Step 2  gh pr merge 4 -R consolelabs/content --squash  (the post)      Exit: 0
        -> content main now has blog/console-log-2026-06.md            verified
Step 3  gh workflow run dispatch.yml -R consolelabs/log.console.so      Exit: 0
        (manual; content's auto-dispatch CONSOLE_PAT is stale)
Step 4  Update submodules (CI) -> commit submodule bump -> Deploy Hugo  conclusion: success

Live checks (curl):
  https://log.console.so/console-log-2026-06/   -> 200   <title>Console Log: June 2026 | Console Labs</title>
  og:description on that page                    -> "The Console Labs lab journal: ..." (metadata fix live)
  https://log.console.so/                        -> 200   footer "© 2026 Console Labs"
  https://log.console.so/log/ , /shape-up/ , /blog/  -> 200 (no regression)
Verdict: PASS
```

## Negative control / rollback

```
BEFORE publish: https://log.console.so/console-log-2026-06/ -> 404 (post not in content main)  [RED]
AFTER  publish: same URL -> 200                                                                  [GREEN]
ROLLBACK (if ever needed): revert content PR #4 on consolelabs/content main (remove
  blog/console-log-2026-06.md), then `gh workflow run dispatch.yml -R consolelabs/log.console.so`;
  the submodule bump drops the post and the URL returns to 404. Reversible.
Verdict: PASS  (the post's liveness is gated entirely on its presence in the content
        vault + a deploy; publishing flips 404 -> 200, reverting flips it back.)
```

## Reproduce / re-verify

```
curl -o /dev/null -w '%{http_code}\n' https://log.console.so/console-log-2026-06/   # 200
curl -s https://log.console.so/ | grep -o '© 20..'                                  # © 2026
```

## Done-gate mapping

- A real monthly post is live: yes, `/console-log-2026-06/` returns 200 with the June report.
- Han-accepted: yes, Han authorized publish 2026-06-28; sign-off block filled in `docs/proof/05-first-post.md`.
- Known caveat (filed): the publish needed a manual dispatch because `CONSOLE_PAT` (2023) is stale; auto-deploy rotation is a Proposed addition.
