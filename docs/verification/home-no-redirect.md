# Verification: home-no-redirect

Behavioral change: the homepage (`layouts/_default/index.html`) JS-redirected `/` to the
log post's URL (`/log/`). Han wants the home to stay at the bare root `/`. Remove the
redirect; the template already renders `blog/log.md` inline via `GetPage`, so the home
now shows the log AT `/` with no bounce.

## Green run

```
Command: env -u GOROOT obsidian-export ./vault ./content && env -u GOROOT hugo -DEF
Exit:    0
Built /index.html: NO window.location.replace; renders the log inline (25,455 bytes,
        contains "A list of changes and updates" = log.md body).
Verdict: PASS
```

## Negative control (revert -> RED -> restore)

```
GREEN   (no script):        / = bare-root, log rendered inline      (no redirect)
REVERT  (re-add script):    / = window.location.replace("/log/")    (RED, bounces away)
RESTORE (no script):        / = bare-root, log rendered inline       (GREEN)
Verdict: PASS  (the redirect script is what bounced the home to /log/; removing it
        leaves the home at the bare root rendering the log.)
```

## Live verification (after merge + deploy)

- [ ] `curl -s https://log.console.so/ | grep -c location.replace` -> 0 (no redirect)
- [ ] `curl -s https://log.console.so/ | grep -o 'A list of changes'` -> present (log at root)

## Note (follow-up, not in this change)

`blog/log.md` still also renders a standalone `/log/` page, so the log content currently
appears at BOTH `/` (home) and `/log/`. Eliminating the `/log/` duplicate (redirect
`/log/` -> `/`, make the log live only at the root) is a content-repo restructure
(log.md `_build.render`, relocate aliases to `_index.md`, check the nav menu) , tracked
separately, pending Han's call.

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
B=$(mktemp -d); env -u GOROOT hugo --destination "$B" -DEF
grep -c location.replace "$B/index.html"   # -> 0
```
