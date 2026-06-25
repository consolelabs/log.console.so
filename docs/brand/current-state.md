# Current-state audit, Console Labs web surfaces

Baseline captured before the 06 design pass. The two public surfaces today are
log.console.so (this repo, the monthly journal) and sticker.console.so (the
kawaii Neko sticker shop). The design work elevates the log without breaking the
shared DNA: white paper, black ink, one mint-green accent, a small Neko mark.

## log.console.so (this repo)

Source of truth: `assets/scss/styles.scss` + `layouts/`. Real tokens pulled from
the SCSS, not eyeballed.

### Tokens in the live skin

| Thing | Current value | Note |
|---|---|---|
| Accent | `--primary-color: #45f1a6` | mint green; the one brand color |
| Accent (light-mode headings) | `--primary-color-light: #2ea370` | darker mint via `prefers-color-scheme: light` |
| Ink | `black` (`#000`) | body text, hardcoded |
| Paper | `white` (`#fff`) | background, hardcoded |
| Sidebar width | `--nav-sidebar-width: 200px` | left nav column |
| Main padding | `--main-padding: 45px` | uniform page padding |
| Body font | `16px sans-serif` | system sans, no webfont, line-height 25px |
| Headings | `sans-serif`, colored `--primary-color` | h1 is `display:none` (title hidden) |
| Mono | `monospace` | `pre` + `code` only |
| Code chip | `background: black; color: white` | inverted terminal chip, bold |
| Tables | `1px solid` black border, collapsed | no zebra; `tabular-nums` on nav |
| Links | `color: black` (white in dark) | underline only on list + ToC links |
| Radius | `2px` (code/pre), `4px` (popover) | minimal rounding |
| Dark mode | `prefers-color-scheme: dark` | `#000` bg, white ink, inverted code + logo |

### Read of the current skin

- **Minimal to a fault.** No webfonts, no type scale, a single accent applied
  bluntly (every heading is mint, which dilutes the accent). h1 is hidden, so
  posts open without a visible title.
- **Already finance-adjacent.** `tabular-nums` shows up on nav; tables are plain
  bordered grids. The instinct is right, it just is not systematized.
- **Layout bones are good.** Fixed left sidebar + content column, collapses under
  1100px. The reskin keeps this structure.
- **The accent is the brand.** `#45f1a6` mint is the one thing carried across both
  surfaces. It stays the anchor.

### What the reskin keeps vs changes

- **Keep:** white/black/mint DNA, the left-sidebar + content-column layout, the
  mobile collapse, `tabular-nums`, the monochrome Neko mark.
- **Change:** add a real type scale + webfonts, demote the accent from "every
  heading" to structural-only (eyebrows, active nav, rules), show the post title,
  soften the code chip from inverted-terminal to a light research-letter surface,
  systematize spacing into an 8px grid.

## sticker.console.so (sibling surface)

The kawaii commerce surface: the colored (yellow) Neko, playful and product-led.
It is NOT in scope to redesign. It sets the shared-brand boundary: the colored
cat and high-saturation play live there. The log is the lab's quieter, ink
sibling, monochrome Neko, editorial restraint. Coherent, not identical: same
mint accent and same Neko character, different register (shop vs journal).

## Visual captures

`docs/brand/visuals/current/` holds the rendered capture of the current log skin.
The live log.console.so was the reference for the "before" baseline that sub-goal
07's before/after proof compares against.
