# Verification: log-revival-slug-basename

Behavioral change: switch the blog permalink rule from `:slug` to `:contentbasename`.
Follow-up defect fix to sub-goal 08: `:slug` derives from the post TITLE, so the
dated monthly file `console-log-2026-06.md` (title "Console Log: June 2026") resolved
at `/console-log-june-2026/`, not the dated `/console-log-2026-06/` that 08's proof and
the publish runbook documented. 08's test only covered `log.md`/`shape-up.md`, where
title == filename, so it missed this. `:contentbasename` keys the URL off the filename,
restoring the dated, sortable, documented path. A post can still override via
front-matter `slug:`.

## Green run

```
Command: env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content && env -u GOROOT hugo list all
Exit:    0
Verdict: PASS

Post URLs (hugo list all, RelPermalink):
  blog/console-log-2026-06.md -> /console-log-2026-06/   (dated, matches filename + docs)
  blog/shape-up.md            -> /shape-up/              (no regression)
  blog/log.md                 -> /log/                   (no regression)
Verdict: PASS

Live local preview (hugo server):
  /console-log-2026-06/ -> 200   /log/ -> 200   /shape-up/ -> 200   /blog/ -> 200
Verdict: PASS
```

## Negative control (revert -> RED -> restore)

```
GREEN   (:contentbasename): blog/console-log-2026-06.md -> /console-log-2026-06/   (dated)
REVERT  (:slug):            blog/console-log-2026-06.md -> /console-log-june-2026/  (title-derived = the bug)
RESTORE (:contentbasename): blog/console-log-2026-06.md -> /console-log-2026-06/   (dated)
Verdict: PASS  (the rule is load-bearing: :slug reintroduces the undated title URL;
        :contentbasename yields the dated filename URL. Existing posts unaffected
        because their title == filename.)
```

## Live verification (after merge + deploy)

Filled once this merges and gh-pages redeploys:

- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/console-log-2026-06/` -> 200 (only once the post itself is published via content PR #4)
- [x] `https://log.console.so/log/` -> 200 (existing post, no regression) , verified live 2026-06-28
- [x] `https://log.console.so/shape-up/` -> 200 , verified live 2026-06-28

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
env -u GOROOT hugo list all | grep console-log-2026-06   # -> /console-log-2026-06/
# control: sd ':contentbasename' ':slug' hugo.yaml; re-run -> /console-log-june-2026/ (RED)
#          sd ':slug' ':contentbasename' hugo.yaml -> restored
```
