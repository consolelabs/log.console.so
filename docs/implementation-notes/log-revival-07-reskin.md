# Implementation notes, sub-goal 07 (reskin)

Deltas from `_meta/megagoals/log-revival/goals/07-reskin-log.md` + `docs/brand/DESIGN.md`.
The spec/DESIGN held the design; this records what they did not pin down.

## 2026-06-25 dark-default via [data-theme], not prefers-color-scheme

- Context: DESIGN.md says "dark-first default, light is the toggle (optional)".
- Decision: dark lives in `:root` unconditionally; the light "day mode" tokens sit
  under `:root[data-theme="light"]`, an explicit opt-in.
- Why: gating light behind `@media (prefers-color-scheme: light)` (the first cut)
  made the default OS-driven, so a light-OS visitor got day mode and never saw the
  brand's signature dark console. `[data-theme]` keeps dark the unconditional default
  and leaves the toggle a one-liner for later (no JS shipped now, per minimum-infra).
- Impact: the brand hero is always dark; a future tiny toggle just sets the attribute.

## 2026-06-25 proof screenshots: Playwright, not `chrome --headless`

- `Google Chrome --headless[=new] --window-size=390` rendered the mobile page at a
  WIDER layout viewport and clipped the image, which looked like a CSS overflow bug
  and sent me chasing a non-existent layout fault (even a `100vw` pin "failed").
- Ground truth via Playwright `getBoundingClientRect` at vw=375: `scrollWidth ==
  clientWidth`, 0 elements past the viewport. The layout was correct the whole time.
- Decision: capture all reskin proof via Playwright (correct mobile viewport via CDP
  device metrics); reverted the unnecessary `100vw` pin back to clean `width:auto`.
- Constraint the spec did not mention: the headless-CLI screenshotter is unreliable
  for responsive proof; trust a measured `scrollWidth` over a screenshot's apparent clip.

## 2026-06-25 layout: block flow on mobile, not flex

- The desktop shell is `header` + a flex row (`nav.menu` sidebar + `main`). On mobile
  (<880px) the row switches to plain `display:block` (sidebar stacks above content)
  rather than a flex/grid collapse. Block flow can't inflate the page width the way a
  flex cross-axis or an `auto-fill minmax` grid track can, so it is the robust choice.

## 2026-06-25 pre-existing warnings left alone (surgical)

- The build emits a `languages.en.languageCode` deprecation (config, owned by 01) and
  `content-bak` duplicate-menu warnings (a stale backup dir inside `content/`). Both
  pre-date this change and are out of the 07 scope (layouts + assets only). Flagged,
  not fixed. The footer's pre-existing em-dash in the copyright line likewise left as-is.
