# Publish runbook, the monthly Console Log

Publishing a month is three moves: run one command, review one PR, click merge.
No manual submodule surgery, no copy-paste into Obsidian, no remembering a dispatch
step. The tooling can never push straight to the live site.

## The flow

```
logwriter  ->  draft (console-log-YYYY-MM.md)
publish    ->  branch + PR in consolelabs/content   [STOPS here]
Han        ->  review + merge the content PR
Actions    ->  content main -> submodule dispatch -> log.console.so gh-pages (live)
```

## Step 1, draft (you)

```
cd ~/workspace/consolelabs/log.console.so/tools/logwriter
./logwriter 2026-06 --out logwriter-out
# writes logwriter-out/console-log-2026-06.md, a vault-ready in-voice draft.
# LLM backend: default `api` (op:// Anthropic key); use `--backend claude` if API
# egress is blocked. Never a hardcoded key.
```

Edit the draft to taste. It is a draft; the tool assembles, you own the voice.

## Step 2, stage + open the content PR (you, one command)

```
./publish logwriter-out/console-log-2026-06.md --dry-run   # optional: preview the plan
./publish logwriter-out/console-log-2026-06.md             # stage + open the PR, then STOP
```

`publish` clones `consolelabs/content`, puts the post at `blog/console-log-YYYY-MM.md`
on a branch `log/YYYY-MM`, opens a PR there, and stops. It prints the PR URL. Nothing
is on `main`; nothing is live.

Flags: `--month YYYY-MM` (override the auto-derived month), `--content-repo owner/name`
(default `consolelabs/content`), `--dry-run` (plan only, no push/PR).

## Step 3, review + merge (Han, the gate)

Open the content PR URL, read the rendered post, merge when happy. Merging is the
only thing that publishes. On merge, the existing GitHub Actions pipeline in the
content repo dispatches to log.console.so, which rebuilds and pushes `gh-pages`. The
post is live in a few minutes.

## Safety properties

- **Draft-only.** `publish` opens a PR; it never pushes to content `main` or to the
  log's `gh-pages`. Verified: after a real run the post exists only on the PR branch,
  not on `main` (`docs/proof/03-publish.md`).
- **Host-portable.** `publish` bakes in no gh-pages internals (just git + a content-
  repo PR), so it survives the planned Cloudflare move untouched; only the deploy in
  Step 3 changes then.
- **No secrets in the tool.** `logwriter` reads its LLM key from `op://` or the
  `claude` CLI at run time. `publish` uses your `gh` auth. The deploy's `GH_PAT` lives
  only in the content repo's Actions secrets, never in this tooling.

## If the deploy does not fire (Step 3)

The dispatch/gh-pages chain needs the content repo's `GH_PAT` Actions secret. If a
merge does not deploy, check the content repo's Actions run + that `GH_PAT` is present
and unexpired. Do not mint or rotate it from the loop; flag it for Han. (Note: content's
own dispatch step uses a SEPARATE `CONSOLE_PAT`; if that one is stale the deploy can be
fired manually with `gh workflow run dispatch.yml -R consolelabs/log.console.so`.)

## Post URLs (and changing them later)

Posts resolve at a root slug derived from the FILENAME: `blog/console-log-2026-06.md`
-> `https://log.console.so/console-log-2026-06/` (set by `permalinks: blog:
/:contentbasename/` in `hugo.yaml`). The `/blog/` section archive list stays at `/blog/`.

To give one post a different URL, add `slug: my-custom-path` to its front-matter (it
overrides the filename rule for that post), or rename the file. **If the post is already
published, changing its slug breaks the old URL** , add an alias so Hugo emits a redirect
from the old path:

```yaml
---
title: "Console Log: June 2026"
slug: june-2026
aliases: ["/console-log-2026-06/"]   # old URL -> redirects to the new one
---
```
