# Verification: log-revival-01-build

Behavioral change: revive the dormant Hugo build on the latest toolchain. Proof
that the build is green now, with a negative control proving the fix is load-bearing.
Companion detail + screenshots: `docs/proof/01-build.md`.

## Green run

```
Command: obsidian-export --hard-linebreaks ./vault ./content
Exit:    0   (56 markdown files produced)
Verdict: PASS

Command: hugo -DEF --minify
Exit:    0   (127 pages / 219 public files; 0 ERROR lines)
Verdict: PASS

Command: hugo server + curl -w '%{http_code}' http://127.0.0.1:1313/  and  /purpose/
Exit:    200 and 200
Verdict: PASS  (rendered, see docs/proof/01-build-home.png + 01-build-post.png)
```

## Negative control (revert -> RED -> restore)

Reverting the load-bearing template fix (`site.Home` back to the removed
`.Sites.First.Home` in `layouts/partials/menu.html`) must turn the build RED.

```
Step 1 GREEN:   hugo -DEF --minify           -> Exit 0, 0 ERROR lines
Step 2 REVERT:  menu.html site.Home -> .Sites.First.Home
        RED:    hugo -DEF --minify           -> Exit 1
                ERROR ... menu.html:9:33 ... <.Sites.First.Home>:
                can't evaluate field First in type page.Sites
Step 3 RESTORE: menu.html .Sites.First.Home -> site.Home (4x)
        GREEN:  hugo -DEF --minify           -> Exit 0, 0 ERROR lines
Verdict: PASS  (the fix is load-bearing: without it the build fails with the
         exact pre-fix error; with it the build is green)
```

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
obsidian-export --hard-linebreaks ./vault ./content
hugo -DEF --minify        # exit 0, writes public/
# negative control:
perl -0pi -e 's/site\.Home/.Sites.First.Home/g' layouts/partials/menu.html
hugo -DEF --minify        # exit 1, "can't evaluate field First"
git checkout layouts/partials/menu.html   # restore
hugo -DEF --minify        # exit 0
```
