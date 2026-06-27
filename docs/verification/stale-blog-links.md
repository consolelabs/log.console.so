# Verification: stale-blog-links

Behavioral change: the homepage (`layouts/_default/index.html`) did
`window.location.replace(currentUrl + "/blog/log")` , a hardcoded redirect to the
legacy `/blog/log` URL. After the permalink change (08 + slug-basename) posts moved off
the `/blog/` prefix to a root slug, so `/blog/log` now 404s and the homepage bounced
visitors to a dead page. Fix: redirect to the log post's real Hugo permalink
(`(site.GetPage "/blog/log").RelPermalink` -> `/log/`), which follows the page wherever
it lives and survives future permalink changes. (`GetPage "/blog/log"` keys off the
CONTENT path, which is unchanged, so the lookup still resolves log.md.)

## Green run

```
Command: env -u GOROOT obsidian-export ./vault ./content && env -u GOROOT hugo -DEF
Exit:    0 (no ERROR lines)
Built homepage redirect: window.location.replace("/log/")   (was the dead /blog/log)
  /log/ built:      yes
  /blog/log/ built: no   (the target the old code pointed at does not exist)
Verdict: PASS
```

## Negative control (revert -> RED -> restore)

```
GREEN   (RelPermalink):       homepage -> location.replace("/log/")           (live)
REVERT  (hardcoded /blog/log): homepage -> location.replace(currentUrl + "/blog/log")  (-> 404, RED)
RESTORE (RelPermalink):       homepage -> location.replace("/log/")           (live)
Verdict: PASS  (the change is load-bearing: the old line targets the dead /blog/log;
        the new line resolves to the post's live permalink.)
```

## Scope

The only stale hardcoded post URL in `layouts/` was this homepage redirect (grep of
`layouts/**` for `/blog/` + hardcoded paths). Post nav/menu links are Hugo-generated
from page permalinks, so they auto-updated. The `/members/<author>` links are a separate
feature, not stale. External / bookmarked links to the old `/blog/<slug>/` URLs are
handled separately by `aliases:` on the moved posts in the content repo (see that PR).

## Live verification (after merge + deploy)

- [ ] `curl -s https://log.console.so/ | grep -o 'location.replace("[^"]*")'` -> `/log/`

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
B=$(mktemp -d); env -u GOROOT hugo --destination "$B" -DEF
grep -o 'location.replace("[^"]*")' "$B/index.html"   # -> /log/
```
