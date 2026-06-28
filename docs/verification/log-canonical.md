# Verification: log-canonical

Behavioral change (#2 of the home/`/log/` fix): the running log renders at BOTH `/`
(home) and `/log/` (its own permalink), a duplicate. Full elimination of `/log/` is
blocked by the toolchain (see the de-obsidian note / #3). As the cheap neutralizer,
point `/log/`'s canonical URL at `/` so search engines treat the home as the real URL
and don't penalize the duplicate. Done via a reusable front-matter override in
`header.html` (`$canonical`, also applied to og:url + twitter:url) + `canonical: /` on
`blog/log.md`.

## Green run

```
Command: env -u GOROOT obsidian-export ./vault ./content && env -u GOROOT hugo -DEF
Exit:    0
Built:
  /log/   <link rel="canonical"> = https://log.console.so/      (points at the home)
  /log/   og:url                 = https://log.console.so/      (matches canonical)
  /       <link rel="canonical"> = https://log.console.so/      (unchanged)
Verdict: PASS
```

## Negative control (cross-page: a page WITHOUT the override keeps its own URL)

```
override page  /log/       (canonical: / )  -> canonical = https://log.console.so/          [GREEN]
control page   /shape-up/  (no override)    -> canonical = https://log.console.so/shape-up/  [default]
control page   /console-log-2026-06/ (none) -> canonical = .../console-log-2026-06/          [default]
Verdict: PASS  (the front-matter `canonical` is load-bearing: only the page that sets it
        gets a redirected canonical; every other page falls back to its own .Permalink.
        So the template change is inert for all pages except the one that opts in.)
```

## Live verification (after merge + deploy)

- [ ] `curl -s https://log.console.so/log/ | grep -o 'rel="canonical" href="[^"]*"'` -> `https://log.console.so/`
- [ ] `curl -s https://log.console.so/shape-up/ | grep canonical` -> `.../shape-up/` (no regression)

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
mkdir -p content; env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
B=$(mktemp -d); env -u GOROOT hugo --destination "$B" -DEF
grep -o 'rel="canonical" href="[^"]*"' "$B/log/index.html"   # -> https://log.console.so/
```

## Note

`/log/` still resolves (200) and still duplicates the home content; this change only
removes the SEO duplicate-content signal. Actually removing `/log/` is tracked as #3 in
the de-obsidian follow-up (NOTES Proposed additions): a build step that seeds
`content/_index.md` from the log so the home owns the content and `log.md` can stop
rendering.
