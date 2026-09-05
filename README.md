# log.console.so

**This repository is archived. The site moved to `consolelabs/console-apps`, `apps/log`.**

`log.console.so` answered from GitHub Pages here until 2026-09-04. It now answers from the Cloudflare Worker `cl-log`, deployed from `consolelabs/console-apps`.

| Fact | Value |
|---|---|
| Cut | `2026-09-04T11:10:50Z` |
| Serves from | Cloudflare Worker `cl-log`, Console Labs account |
| Source now | `consolelabs/console-apps`, `apps/log` |
| Soak verdict | PASS, 2026-09-05 |
| GitHub Pages here | disabled 2026-09-05 |

The full history of this repository was imported into `consolelabs/console-apps` under `apps/log`, so `git log --follow` there reaches the first commit made here. Nothing was lost by archiving this copy.

## Where the posts live, and how one gets published

The posts were never in this repository. They live in `consolelabs/content`, which this repository pulled in as the `vault` submodule and which is NOT archived.

A merge to `main` in `consolelabs/content` now dispatches the `Publish log` workflow in `consolelabs/console-apps`, which builds the site against `consolelabs/content` `main` and deploys `cl-log`. Nothing polls, nothing is scheduled, and no submodule pin has to be remembered. The old hourly build that ran here is gone.

| Record | Where |
|---|---|
| The publish chain, end to end | `consolelabs/console-apps` `docs/publish-pipeline.md` |
| The port | `consolelabs/console-apps` `docs/verification/log-port.md` |
| The cut, with every command and its exit code | `consolelabs/console-apps` `docs/verification/log-cut.md` |
| The runbook the cut followed | `consolelabs/console-apps` `docs/cut-runbook-log.md` |
| The migration program | `tieubao/console-labs` `docs/specs/SPEC-013-web-migration-vercel-to-cloudflare.md` |

To write a post, send it to `consolelabs/content`. To change the site itself, send it to `consolelabs/console-apps` `apps/log`. Nothing here is deployed.
