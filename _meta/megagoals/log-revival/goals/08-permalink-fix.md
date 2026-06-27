# Sub-goal 08: fix the post permalink slug

**Merge policy:** auto
**Time budget:** 1 calendar day
**Proof:** before/after URL table + a screenshot of a post at its clean URL on log.console.so, in `docs/proof/08-permalink.md`
**Depends on:** 07 (the reskin must be live to verify against the real site) + 03 (a staged post to view)
**Branch:** fix/log-revival-08-permalink
**PR base:** main

## Outcome

A monthly post resolves at a clean, canonical URL instead of the current redundant `log.console.so/blog/log/`. Today the post path doubles the same concept ("blog" + "log") on a site that already IS the log, and `/blog/log/` is actually the section list page (title "Log | Console Labs"), not a per-post slug derived from the filename (`blog/console-log-2026-06.md`). After this sub-goal, each post has its own predictable slug and there is no stale `/blog/log/` duplicate.

## Quality bar

A reader can guess the URL of a month's post and a new post never collides with the section page. The path reads clean for a site that is itself a log (no "blog/log" doubling). Permalink structure is declared in config, not accidental.

## How to close the loop

- Decide the canonical post URL with Han: either drop the redundant `blog/` prefix (e.g. `/2026-06/` or `/console-log-2026-06/`) or keep a single non-doubled segment. Record the decision.
- Implement it via a `permalinks:` block in `hugo.yaml` (and/or a front-matter `slug` on the post), so the section list page and the per-post page no longer collide on `/blog/log/`.
- Rebuild locally (`make build`), confirm the post renders at the new path and the old `/blog/log/` post-dup is gone (the section index may legitimately stay as a list page).
- After the live deploy, screenshot the post at its clean URL into `docs/proof/08-permalink.md` with a before/after URL table.

**Done =** `hugo.yaml` declares the post permalink, a built post resolves at the agreed clean slug, the redundant `/blog/log/` post duplicate is gone, and `docs/proof/08-permalink.md` carries the before/after table + a live screenshot.

## Scope edges

**In:** `hugo.yaml` permalink config, optional front-matter `slug`, verifying the section-vs-post URL split, the before/after proof.
**Out:** the writing tool (02), the publish flow (03), the brand/CSS (06/07), monitoring (04).
**Not:** restructuring the content vault's directory layout, migrating to Cloudflare, changing the post's content, rotating the deploy secret.

## Where to look

`hugo.yaml` (no `permalinks:` block today), the post front-matter in `consolelabs/content` (`blog/console-log-2026-06.md`, no `slug`), the live `/blog/log/` section page, the `outputs.section` config.

## PR body

Fix the redundant post permalink: a monthly post resolved at `/blog/log/` (the section page, "blog" + "log" doubled) instead of a per-post slug. Declares the post permalink in `hugo.yaml` (+ optional front-matter `slug`) so posts get a clean canonical URL and no longer collide with the section index.

Verify: built post resolves at the agreed slug; `/blog/log/` post-dup gone; before/after table + live screenshot in `docs/proof/08-permalink.md`. Part of mega-goal `log-revival`.

## Notes

Raised by Han 2026-06-28 from the live site (screenshot: `log.console.so/blog/log/`). Promoted from NOTES `## Proposed additions` into a formal sub-goal on Han's call.
