# Verification: site-metadata

Behavioral change: refresh site-level metadata in `hugo.yaml` , bump the copyright
year (2023 -> 2026, both the top-level `copyright` and `params.Copyright` used in the
JSON-LD) and add a site-wide `params.description`. The description was UNSET, so every
page with no own front-matter description (including the new monthly posts, whose
`description:` is empty) rendered an EMPTY `<meta name="description">` / `og:description`
/ JSON-LD description. The home page is unaffected (it carries its own description in
`_index.md`).

## Green run

```
Command: env -u GOROOT obsidian-export ./vault ./content && env -u GOROOT hugo -DEF
Exit:    0
Verdict: PASS

Built HTML:
  /console-log-2026-06/  <meta name="description"> = "The Console Labs lab journal:
      monthly notes on what we build, ship, and learn. Safe, accessible tooling,
      narrated honestly."   (was empty before)
  home + footer          "© 2026 Console Labs"   (was "© 2023 ConsoleLabs.")
Verdict: PASS
```

## Negative control (revert -> RED -> restore)

```
GREEN   (params.description set): /console-log-2026-06/ meta description = the site tagline
REVERT  (strip params.description): /console-log-2026-06/ meta description = ""   (EMPTY = RED)
RESTORE (params.description set): meta description = the site tagline again
Verdict: PASS  (the param is load-bearing: without it posts with an empty front-matter
        description ship an empty meta/OG description; with it they inherit the site tagline.)
```

## Live verification (after merge + deploy)

- [ ] `curl -s https://log.console.so/ | grep -o '© 20..'` -> `© 2026`
- [ ] the first post (once published) carries a non-empty og:description

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
B=$(mktemp -d); env -u GOROOT hugo --destination "$B" -DEF
grep -o '<meta name="description"[^>]*>' "$B/console-log-2026-06/index.html"
# control: sd 'description: "The Console.*honestly."' '' hugo.yaml; rebuild -> empty (RED)
#          git checkout -- hugo.yaml -> restored
```
