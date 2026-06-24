# Sub-goal 03: the reviewed publish flow

**Merge policy:** gate
**Time budget:** 1 calendar day after PR-02 merges
**Proof:** run-table of the publish helper (dry-run + the staged result), plus 1-2 screenshots of the post live on log.console.so after Han triggers the deploy
**Depends on:** 02
**Branch:** feat/log-revival-03-publish
**PR base:** feat/log-revival-02-writer

## Outcome

A one-command, draft-only publish flow that takes a drafted monthly post and stages it for release: it lands the post in the `consolelabs/content` vault repo on a branch and opens a PR there, then stops. Merging that content PR + triggering the deploy stays Han's hand (draft-only). When Han merges, the existing GitHub Actions pipeline (submodule dispatch -> main -> gh-pages) publishes it live. Nothing reaches the live site without Han's merge.

## Quality bar

The whole act of publishing a month is: run one command, review one PR, click merge. No manual submodule surgery, no copy-paste into Obsidian, no remembering the dispatch step. Boring and safe: the tool can never push straight to the live site.

## How to close the loop

- Build the publish helper (a thin script/command, host-portable: it operates on the content vault + opens a PR, it does NOT bake in gh-pages internals, since the site moves to Cloudflare later). Capture a dry-run run-table into `docs/proof/03-publish.md`.
- Drive it end to end for a test/sample post: stage into the `consolelabs/content` repo on a branch, open the content PR, STOP. Record the PR URL.
- After Han merges the content PR and the deploy runs, screenshot the post live on log.console.so into `docs/proof/`. This is the gate step (outward + needs the `GH_PAT` deploy secret); the loop prepares + opens, Han merges.
- Document the publish runbook (the one command + the merge step + how the dispatch/deploy chain fires) at `docs/publish-runbook.md`.

**Done =** the publish helper stages a drafted post into `consolelabs/content` and opens a content PR (no direct push to live), the runbook documents the flow, and `docs/proof/03-publish.md` carries the dry-run table + the live screenshot taken after Han's merge.

## Scope edges

**In:** the publish helper, opening the content-repo PR, the publish runbook, verifying the existing deploy chain still lands content live.
**Out:** the writing tool itself (02), monitoring (04), rebuilding the deploy from scratch.
**Not:** full-auto publish (Han chose draft-only), pushing directly to gh-pages, migrating to Cloudflare now, provisioning/rotating the GH_PAT secret (flag if expired, do not mint), bypassing the content-repo PR.

## Where to look

The `consolelabs/content` repo (the `vault/` submodule), the existing `.github/workflows/{main,dispatch}.yml`, the `logwriter` output from 02, `docs/voice.md`.

## PR body

A draft-only, one-command publish flow: stage a drafted monthly post into the `consolelabs/content` vault repo and open a PR there; Han merges to publish via the existing Actions deploy. Host-portable (no gh-pages internals baked in, ahead of the Cloudflare move).

Verify: run the publish helper dry-run; it opens a content-repo PR without touching the live site; after merge the post is live (screenshot). See `docs/proof/03-publish.md` + `docs/publish-runbook.md`.

Stacked on #<PR-02>; review after it. GATE: opens the content PR then stops for Han's merge. Part of mega-goal `log-revival`.

## Notes
