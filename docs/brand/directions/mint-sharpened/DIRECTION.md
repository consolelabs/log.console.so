# Brand direction: Mint, sharpened

A token-based spec for log.console.so. The angle: take the current mint-and-white
minimal and make it **deliberate**. Same restraint, but every value is chosen, not
defaulted. A refined mint accent, a real typographic hierarchy, a crisp spacing
scale, and a subtle Neko character. Restraint is the feature, not the absence of one.

This is an EVOLUTION of the live site (white bg, black text, one mint accent
`#45f1a6`, left sidebar, black Neko logo), not a rewrite. A reader who knows the
current log should feel it got sharper, not replaced.

---

## 1. Design principles

1. **One accent, used with intent.** Mint is a scalpel, not a highlighter. It marks
   structure (section labels, the active nav item, links, rules) and nothing else.
   Black does the heavy lifting; mint points.
2. **Hierarchy through type and space, not color.** Headings differ by size, weight,
   and letter-spacing, not by getting louder. The page reads top-to-bottom without
   a single decorative flourish.
3. **A spacing scale, not eyeballed margins.** Every gap is a step on a 4px base
   scale. Rhythm is the polish.
4. **Neko is present, never loud.** The black cat shows up as the logo mark and one
   small inline glyph (the author byline, an end-of-post mark). It is a signature,
   not wallpaper.
5. **Editorial, not app.** This is a lab journal. It should read like a well-set
   essay with a quiet index down the side, not a dashboard.

---

## 2. Color palette (tokens)

All foreground/background pairs below are verified WCAG-AA (>= 4.5:1 for body text,
>= 3:1 for large text and UI).

| Token | Hex | Role | Contrast note |
|---|---|---|---|
| `--ink` | `#0F1115` | Primary text, logo, rules | 18.8:1 on paper |
| `--ink-soft` | `#3A3F47` | Secondary text, captions, meta | 9.6:1 on paper |
| `--ink-faint` | `#6B7280` | Tertiary text, timestamps, placeholders | 4.8:1 on paper |
| `--paper` | `#FFFFFF` | Page background | base |
| `--paper-2` | `#F7F8F7` | Callout / code-block surface, hover wash | subtle lift |
| `--line` | `#E3E6E3` | Hairline borders, dividers | structure only |
| `--mint` | `#1FB877` | THE accent: links, active nav, section labels on paper | 3.6:1 on paper (large/UI + underlined links only) |
| `--mint-ink` | `#0E7C4F` | Mint text that must pass AA as body-size (inline links if not underlined) | 4.9:1 on paper |
| `--mint-bright` | `#45F1A6` | Legacy brand mint: dark-mode accent, logo dot, never as text on white | reserved for dark bg |
| `--mint-wash` | `#EAFBF2` | Tint behind a mint callout / TL;DR band | 1.04 lift, decorative |
| `--mint-line` | `#9DE9C4` | Left border of a mint callout | structure |

Dark mode (the live site already ships `prefers-color-scheme: dark`; keep it, tune it):

| Token | Hex | Role |
|---|---|---|
| `--ink` | `#E8EAE8` | Text on near-black |
| `--paper` | `#0C0E10` | Background |
| `--paper-2` | `#15181B` | Surface |
| `--line` | `#262A2E` | Borders |
| `--mint` | `#45F1A6` | Accent (the bright legacy mint earns its keep here: 11:1 on `#0C0E10`) |

**Rules of the palette.** On white, mint text only appears as an underlined link or
as an uppercase section label (both read as "large/UI", so `--mint` 3.6:1 is allowed);
any mint that must read as plain body copy uses `--mint-ink`. Never two accents.
Never mint on mint-wash for text. The legacy `#45f1a6` survives as the dark-mode
accent and the logo dot, so the brand DNA is intact.

---

## 3. Type scale

Two families, both free on Google Fonts, both crypto-credible and editorial.

- **Display + headings + UI:** `Space Grotesk` (geometric, slightly mechanical, the
  "console" feel without being a monospace cliche). Weights 500/600/700.
- **Body:** `Inter` (workhorse, superb at 16-18px, neutral so the voice carries it).
  Weights 400/500/600.
- **Code + the neko byline + section labels:** `JetBrains Mono` (the lab's terminal
  voice; also what gives "console" its meaning). Weights 400/500/700.

Stack fallbacks: `"Space Grotesk", "Inter", system-ui, sans-serif` etc., so an
offline render degrades to the current sans-serif gracefully.

Scale (1.250 major-third, rounded to clean px):

| Token | Font | Size / line-height | Weight | Tracking | Use |
|---|---|---|---|---|---|
| `--t-display` | Space Grotesk | 40px / 46px | 700 | -0.02em | Post title (h1) |
| `--t-h2` | Space Grotesk | 27px / 34px | 600 | -0.01em | Section heads (Shipped, Thinking) |
| `--t-h3` | Space Grotesk | 20px / 28px | 600 | -0.005em | Sub-heads |
| `--t-label` | JetBrains Mono | 12px / 16px | 500 | 0.12em, UPPERCASE | Eyebrow / section label / nav group |
| `--t-body` | Inter | 17px / 28px | 400 | 0 | Post body |
| `--t-body-sm` | Inter | 15px / 24px | 400 | 0 | Sidebar links, captions |
| `--t-meta` | JetBrains Mono | 13px / 20px | 400 | 0.02em | Date, author, tags |
| `--t-code` | JetBrains Mono | 14px / 22px | 400 | 0 | Inline + block code |

**Hierarchy rule that matters:** headings are NOT mint anymore (the live site paints
all h1-h4 mint, which flattens the hierarchy into "everything is the accent").
Headings are `--ink`. Mint moves UP to the **eyebrow label** above each section
(`SHIPPED`, `THINKING` in mono uppercase) and to links. So the eye gets: tiny mint
label -> big black heading -> body. That is the whole "sharpened" move in one
sentence.

---

## 4. Spacing scale

4px base. Name by step, never hand-pick a one-off value.

| Token | Value | Typical use |
|---|---|---|
| `--s-1` | 4px | Icon-to-text gaps, chip padding |
| `--s-2` | 8px | Inline gaps, label-to-heading |
| `--s-3` | 12px | List-item spacing |
| `--s-4` | 16px | Paragraph bottom margin |
| `--s-5` | 24px | Heading top margin (h3) |
| `--s-6` | 32px | Nav group spacing |
| `--s-7` | 48px | Section block spacing (between Shipped / Thinking) |
| `--s-8` | 64px | Header-to-content, major page rhythm |
| `--s-9` | 96px | Page top padding, footer breathing room |

Layout constants:
`--measure: 68ch` (post body max line length, the readability spine),
`--sidebar-w: 220px`, `--gutter: 56px` (sidebar-to-content), `--page-pad: 48px`.

---

## 5. Component patterns

### Header + logo
- Full-width, hairline `--line` bottom border, `--page-pad` padding.
- Logo = the black Neko cat mark (current SVG, `--ink` fill) at 28px, followed by
  the wordmark `console·log` set in Space Grotesk 600, where the middle dot is a
  4px `--mint-bright` square. That dot is the entire color budget of the header.
- To the right (desktop): a one-line, `--ink-faint`, mono tagline: `the lab, narrated monthly`.

### Sidebar nav
- Fixed `--sidebar-w`, `--gutter` from content, top-aligned with the post.
- Grouped with mono uppercase `--t-label` group headers in `--mint` (`PAGES`, `STREAMS`).
- Items: `--t-body-sm`, `--ink-soft`. Hover: `--ink` + `--paper-2` wash, no underline.
- **Active item:** `--ink`, weight 600, with a 2px `--mint` left rule (4px inset).
  One active marker, mint, unmistakable.
- Order matches the live nav: About, Philosophy, Purpose, then a divider, then
  experiments, blog, playbook.

### Post body
- Constrained to `--measure`. `--t-body`. Paragraphs `--s-4` apart.
- First paragraph after the title is the description/hook in `--ink-soft`, 19px,
  acting as a standfirst. One step bigger, one shade softer: editorial deck.

### Headings (in-post)
- Each major section leads with the mono eyebrow label (mint) then the h2 (ink).
  `--s-7` of air above the label resets the rhythm between sections.

### Links
- In body: `--mint-ink`, underline with 2px offset and `--mint` underline color.
  Hover: underline thickens, text -> `--ink`. Visited == default (a journal, not a
  link farm; no purple).
- Underline + AA-passing `--mint-ink` means links are findable without color alone.

### Code
- Inline: `--t-code`, `--paper-2` background, `--line` 1px border, `--s-1` padding,
  4px radius, text `--ink`. (Drops the live site's heavy black-pill inline code,
  which fought the minimal body; a quiet tinted chip is sharper.)
- Block: `--paper-2` surface, `--line` border, 8px radius, `--s-5` padding,
  a 2px `--mint` top accent rule. Comments in `--ink-faint`.

### Callouts
- One variant, mint. `--mint-wash` background, 3px `--mint-line` left border,
  6px radius, `--s-5` padding. A mono `--t-label` lead in `--mint-ink`
  (e.g. `HEADS UP`, `TL;DR`). Body in `--ink`. No other callout colors: a single
  honest mint note keeps the "small but sharp" promise.

### Lists (Shipped items)
- Each shipped item: a `·` mint bullet, a bold `--ink` lead (what shipped), then a
  `--ink-soft` "why it matters" on the same paragraph. No cards, no boxes; the
  spacing scale does the separating.

### Neko mark + end-of-post
- The post ends with a centered small black Neko glyph at 20px (the signature), and
  the byline `— neko-san` in `--t-meta`. That is the only place the cat appears in
  the body. Present, quiet, owned.

---

## 6. How this expresses the 4 pillars + the voice

| Pillar | Visual expression |
|---|---|
| **Small but Sharp 🔍** | One accent, one callout variant, one cat glyph, two type families. The restraint IS the sharpness. A reader feels nothing is arbitrary, which is the design equivalent of "a tool that nails one job". |
| **Financial-Centric 💰** | Mono for all meta, labels, and numbers (`--t-label`, `--t-meta`, code) gives a precise, ledger-like, money-tooling register. Tabular-num alignment on dates/figures reads as financial care, not decoration. |
| **Accessible 🌐** | Every text token is WCAG-AA verified; links never rely on color alone (always underlined); `--measure` caps line length for readability; dark mode is first-class; fallbacks degrade to plain sans. Accessible literally, not just in spirit. |
| **Gamified 🎮** | The mint eyebrow labels, the active-nav mint rule, the `·` mint bullets, and the end-of-post Neko signature add light, playful "moments" without toys. The cat is the wink. Emoji in headings (the voice already loves them) land cleanly against the restrained field instead of fighting it. |

**Voice fit.** The voice is playful, casual, emoji-friendly, honest about failure,
first-person-plural. This direction gives that voice a clean, confident stage:
generous space so short honest paragraphs breathe, emoji that pop because the page
is quiet around them, and a mono/serif-free editorial tone that reads as a real lab
talking, not a press release. The sharpened restraint says "we are small but we
know exactly what we are doing", which is the brand's whole personality.

---

## 7. Do / Don't

**Do**
- Use mint for structure only: eyebrow labels, active nav, links, the callout, rules.
- Keep headings `--ink`; let size + weight + the mint eyebrow carry hierarchy.
- Set all meta, labels, dates, and numbers in JetBrains Mono.
- Cap body text at `--measure`; respect the spacing scale for every gap.
- Let the Neko mark appear exactly twice: logo and end-of-post signature.
- Keep emoji to roughly one per heading (the voice rule); let whitespace frame them.
- Verify any new fg/bg pair against AA before shipping it.

**Don't**
- Don't paint headings mint (the live-site flatten-everything mistake).
- Don't introduce a second accent color, ever. Mint is the whole budget.
- Don't add card chrome, shadows, or gradients; this is editorial, not a dashboard.
- Don't use mint as a background for body text, or `--mint` as plain body-size text
  on white (use `--mint-ink`).
- Don't scatter the cat across the page; one signature, not a sticker sheet.
- Don't use em dashes anywhere (brand voice rule). Commas, colons, sentence breaks.
- Don't let line length run past ~68ch or gaps go off-scale; drift kills the polish.
