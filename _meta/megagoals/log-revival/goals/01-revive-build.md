# Sub-goal 01: revive the build

**Merge policy:** auto
**Time budget:** 1-3 hours of loop work
**Proof:** run-table (the build command, exit code, real stdout tail) + 1-2 screenshots of the locally-served site rendering a real page
**Depends on:** none
**Branch:** chore/log-revival-01-build
**PR base:** main

## Outcome

The dormant pipeline builds again on today's toolchain. From a clean checkout, the Obsidian vault converts and Hugo produces `public/` with no errors, and the local dev server renders the existing content correctly. The live deploy is NOT touched here (that is sub-goal 03); this is purely "the build is alive again."

## Quality bar

Boring and bulletproof. Someone clones this in a year and `make build` just works, with the toolchain versions pinned so Hugo or obsidian-export drift never silently breaks a publish again.

## How to close the loop

- Run the build path the repo already defines (`make build`, which runs `obsidian-export` then `hugo -DEF --minify`). Capture the command, exit code, and the stdout tail into a run-table at `docs/proof/01-build.md`.
- Fix only what 2 years of bit-rot broke: Hugo version (extended), the `obsidian-export` cargo install, devbox pins, any submodule path assumption. Pin the versions you land on (devbox.lock / a noted hugo version) so the fix is durable.
- Serve locally (`hugo server` via `make run` or `make watch-run`) and screenshot the homepage + one inner post rendering correctly into `docs/proof/`.

**Done =** a clean `make build` exits 0 and writes `public/`, AND `docs/proof/01-build.md` carries the run-table plus the two render screenshots.

## Scope edges

**In:** the build toolchain (Makefile, devbox, hugo config, obsidian-export step), version pinning, the `vault/` submodule wiring.
**Out:** the live gh-pages deploy (sub-goal 03), any content rewriting, the site theme/design.
**Not:** redesigning the layouts, upgrading Hugo to a major new theme, migrating to Cloudflare (a separate future effort), adding new content sections.

## Where to look

The repo root build files (Makefile, devbox.json, hugo.yaml, inotify.sh, scripts/), the `vault/` submodule, the existing `layouts/`.

## PR body

Revive the dormant log.console.so build on the current toolchain (Hugo + obsidian-export + devbox). Local `make build` produces `public/`; versions pinned so publishes stop silently rotting.

Verify: `make build` exits 0 and writes `public/`; see `docs/proof/01-build.md` for the run-table + render screenshots.

Part of mega-goal `log-revival` (see `_meta/megagoals/log-revival/ROADMAP.md`).

## Notes
