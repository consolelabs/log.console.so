# Sub-goal 07: reskin the log site to the new brand

**Merge policy:** gate
**Time budget:** 1 calendar day after PR-06 merges
**Proof:** 2-3 before/after screenshots (homepage, a post, and mobile width) of the log site on the new brand, plus a token-by-token checklist against DESIGN.md
**Depends on:** 06, 01
**Branch:** feat/log-revival-07-reskin
**PR base:** docs/log-revival-06-design

## Outcome

The log.console.so Hugo `layouts/` + `assets/` implement the `docs/brand/DESIGN.md` brand guideline. The live-rendered site (homepage, a post, the list pages, mobile) matches the locked tokens: colors, type, spacing, components. It looks like the new Console Labs brand, coherent with sticker.console.so, not the 2024 skin.

## Quality bar

Looks like the OS did it, not a bolt-on. On-brand, responsive, and clears accessibility basics (text contrast, visible focus states). A reader feels the lab's visual identity the moment the page loads. No half-applied tokens, no leftover old styling.

## How to close the loop

- Implement the DESIGN.md tokens in `layouts/` + `assets/` (CSS variables for color/type/spacing, the component patterns). Keep it Hugo + CSS; no framework rewrite.
- Build + serve locally (`make run` / devbox), capture before/after screenshots for the homepage, one post, and a mobile width into `docs/proof/07-reskin.md`.
- Walk DESIGN.md token by token and tick each against the rendered site in a checklist inside the proof file.

**Done =** the locally-served log site renders on the DESIGN.md tokens (color/type/spacing/components match), AND `docs/proof/07-reskin.md` carries the before/after screenshots (desktop + mobile) and the token checklist.

## Scope edges

**In:** `layouts/` (single, list, index, partials), `assets/` + CSS, the design tokens, responsive + a11y basics.
**Out:** DESIGN.md itself (06), content, the build toolchain (01), deploy + monitoring.
**Not:** a JS framework rewrite (stays Hugo), a CMS, new content sections, redesigning the content structure, the Cloudflare move, re-opening the brand decisions locked in 06.

## Where to look

`layouts/` (single.html, list.html, index.html, partials/), `assets/`, `hugo.yaml`, and `docs/brand/DESIGN.md` (the spec from 06).

## PR body

Reskin log.console.so to the `docs/brand/DESIGN.md` brand guideline: implement the locked color/type/spacing/component tokens in the Hugo layouts + assets. Stays Hugo, responsive, a11y basics.

Verify: the locally-served site matches the DESIGN.md tokens; see `docs/proof/07-reskin.md` for before/after (desktop + mobile) + the token checklist.

Stacked on #<PR-06>; review after it. GATE: Han signs off on the visual result. Part of mega-goal `log-revival`.

## Notes
