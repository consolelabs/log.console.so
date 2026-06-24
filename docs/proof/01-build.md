# Proof of done, sub-goal 01: revive the build

Reviving the dormant (Feb 2024) Hugo build on the latest toolchain, plus the
"upgrade the stack to latest" directive. Templates fixed, versions bumped to
latest, build green.

## Acceptance criteria

- Clean build exits 0 and writes `public/`.
- Build runs on the latest toolchain (not the 2-year-old pins).
- This proof carries a run-table + render screenshots.

## Toolchain (verified)

| Tool | Was (pinned) | Now (verified) |
|---|---|---|
| Hugo | 0.120.2 | 0.163.3 extended (latest) |
| obsidian-export | stale cached binary (broken: missing LC_UUID) | 25.3.0 (rebuilt) |
| devbox pins | fixed versions | `latest` (go, hugo, obsidian, rust, etc.); lock regenerated |
| CI actions | cache@v1, actions-rs@v0.1 (archived), hugo@v2, gh-pages@v3 | cargo-install@v3, hugo@v3, gh-pages@v4 |

## What the version jump broke (and the fix)

| Symptom | Root cause | Fix |
|---|---|---|
| `obsidian-export` failed: `missing LC_UUID load command`, 0 md files | cached cargo binary incompatible with current macOS dyld | `cargo install obsidian-export --force` (rebuilt to 25.3.0) |
| `can't evaluate field First in type page.Sites` (build ERROR) | `.Sites.First.Home` removed in modern Hugo | `site.Home` in `layouts/partials/menu.html` (4x) |
| `can't evaluate field Author in type page.Site` (build ERROR) | top-level `author` config + `.Site.Author` dropped | moved `author` under `params:` in `hugo.yaml`; `site.Params.author.name` |
| deprecation warns: `.Site.LanguageCode`, root `languageCode` | deprecated in Hugo 0.158 | `site.Language.Locale`; `languageCode` moved under `languages.en` |

## Confirmation run-table

| # | Command | Result |
|---|---|---|
| R1 | `obsidian-export --hard-linebreaks ./vault ./content` | exit 0, 56 md files produced |
| R2 | `hugo -DEF --minify` | exit 0, 127 pages / 219 public files, no ERROR (only non-fatal "no json layout" warns) |
| R3 | `hugo server` + `curl -w %{http_code}` homepage | 200; `<title>Console Labs</title>` |
| R4 | `curl` `/purpose/` | 200 |
| R5 | empty-content `hugo -DEF --minify` (template smoke test) | exit 0 (basis for the new `build-check.yml` PR CI) |
| R6 | Playwright render, homepage + `/philosophy/` | 0 console errors; screenshots below |

## Render screenshots

- Homepage (the running monthly log): `docs/proof/01-build-home.png`
- A content page (Philosophy): `docs/proof/01-build-post.png`

Both render the full site: Neko logo, menu, experiments/blog/playbook nav, and
real content (the existing Jan 2024 / Dec 2023 monthly logs).

## Hardening added

- `.github/workflows/build-check.yml`: a `pull_request` smoke test that builds
  templates+config with empty content (no private submodule needed), so future
  PRs catch template/config bit-rot before merge. This is the mega-goal's PR CI.

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
obsidian-export --hard-linebreaks ./vault ./content   # 56 md files
hugo -DEF --minify                                     # exit 0, writes public/
hugo server -DEF --port 1313                            # serve + eyeball
```

## Note

The `languages.en` restructure clears the root `languageCode` deprecation; the
only remaining warns are "no json layout for taxonomy/section/term" (pre-existing,
non-fatal). Filed as a Proposed addition, not a blocker.
