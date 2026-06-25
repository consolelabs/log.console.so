# Brand Direction: Fintech Editorial

> One brand direction for **log.console.so**, the public monthly lab journal of Console Labs.
> Author persona: `neko-san`.

## The angle in one line

A sharp crypto-research publication that happens to have a sense of humor. Serious,
confident typography. A restrained palette with one mint accent. Generous whitespace.
A finance sensibility (mono figures, tabular numbers, hairline rules). The playfulness is
seasoning, an emoji and a wink, never the main course.

We are elevating the current site, not replacing its DNA. Today: white, black, sans, one
mint-green accent (`#45f1a6`), left sidebar, black Neko logo, very minimal. We keep all of
that and make it look like it knows what it is worth.

---

## 1. Design principles

1. **Editorial, not appy.** This reads like a research letter you would forward, not a SaaS
   landing page. Columns are measured, headings carry weight, the eye has room.
2. **Numbers are first-class.** Anything that is a figure, a date, a version, a percentage,
   a hash, gets the mono face and tabular alignment. Finance literacy shows in the details.
3. **One accent, used with discipline.** Mint (`#45f1a6`) is a scalpel, not a highlighter.
   It marks structure (section eyebrows, rules, the active nav, inline emphasis), never body
   text, never large fills.
4. **Whitespace is a feature.** Restraint signals confidence. When in doubt, add space, not ink.
5. **The cat earns its keep.** Neko is the byline and the wordmark, monochrome and small. It is
   a sign-off, not a sticker. The kawaii color version lives on sticker.console.so; here Neko is
   the lab's quiet mascot in ink.

---

## 2. Color palette (tokens)

All pairings below meet WCAG-AA (4.5:1 for body text, 3:1 for large text and UI) on their
stated background.

### Light mode (default)

| Token | Hex | Role | Contrast note |
|---|---|---|---|
| `--ink` | `#0f1115` | Body text, wordmark, primary headings | 18.4:1 on paper |
| `--ink-soft` | `#3a3f47` | Secondary text, captions, metadata | 9.1:1 on paper |
| `--ink-faint` | `#6b727c` | Tertiary text, timestamps, nav-inactive | 4.9:1 on paper |
| `--paper` | `#ffffff` | Primary background | base |
| `--paper-2` | `#f7f8f7` | Sidebar, callout, code surface | subtle lift |
| `--rule` | `#e3e6e4` | Hairlines, table borders, dividers | structural |
| `--mint` | `#45f1a6` | Brand accent: eyebrows, active nav, rules, marks | decorative + structural |
| `--mint-ink` | `#0a7a4a` | Accent *as text* (links, inline emphasis) on paper | 4.6:1 on paper, AA |
| `--mint-wash` | `#eafbf2` | Tint behind pull-quotes / "shipped" callouts | base for ink text |
| `--amber-ink` | `#8a5a00` | "Tinkering / flopped" callout text | 4.7:1 on `--amber-wash` |
| `--amber-wash` | `#fbf3e3` | "Tinkering / honest-about-failure" callout fill | base |

**Key contrast rule:** `--mint` (`#45f1a6`) is a *light* green. It is for fills, rules, and
marks behind dark ink, never for text on white. When mint must read as text (links, inline
code emphasis), use `--mint-ink` (`#0a7a4a`), which clears AA.

### Dark mode (`prefers-color-scheme: dark`)

| Token | Hex | Role |
|---|---|---|
| `--ink` | `#e8ebe9` | Body text |
| `--ink-soft` | `#aab1ab` | Secondary |
| `--ink-faint` | `#7d847e` | Tertiary |
| `--paper` | `#0c0e0d` | Background |
| `--paper-2` | `#151917` | Sidebar / code surface |
| `--rule` | `#262b28` | Hairlines |
| `--mint` | `#45f1a6` | Accent, unchanged (pops on dark) |
| `--mint-ink` | `#5cf3b0` | Accent as text on dark (8.9:1, AA) |
| `--mint-wash` | `#10231a` | Pull-quote / shipped fill |
| `--amber-ink` | `#e8c87a` | Tinkering text on dark |
| `--amber-wash` | `#241d10` | Tinkering fill |

---

## 3. Typography

Two families plus one mono. Pulled from Google Fonts so the mockup is self-contained.

| Role | Family | Fallback | Why |
|---|---|---|---|
| Headings + wordmark | **Fraunces** (opsz, soft serif) | `Georgia, serif` | A modern, slightly witty serif. Confident and editorial without being stuffy. The "tasteful serif for headings" call. |
| Body + UI | **Inter** | `system-ui, sans-serif` | A refined grotesk. Neutral, finance-credible, superb at small sizes and tabular numbers. |
| Figures + code + metadata | **JetBrains Mono** | `ui-monospace, monospace` | Data sensibility. Dates, versions, percentages, hashes, code. Tabular by default. |

> Alternative all-sans take (if a serif feels too soft for a given month): set headings to
> Inter at weight 700 with `-0.02em` tracking. The structure below is unchanged; only the
> heading family swaps. Keep ONE choice site-wide, do not mix.

### Type scale (1.250 major-third, 16px base)

| Token | Use | Size / line-height | Family / weight |
|---|---|---|---|
| `--t-wordmark` | Logo lockup | 22px / 1 | Fraunces 600 |
| `--t-display` | Post title (h1) | 40px / 1.1 | Fraunces 600, `-0.01em` |
| `--t-h2` | Section heads (TL;DR, Shipped) | 26px / 1.2 | Fraunces 600 |
| `--t-h3` | Sub-heads | 19px / 1.3 | Inter 650 |
| `--t-eyebrow` | Section eyebrow / kicker | 12px / 1.2, `0.14em`, uppercase | JetBrains Mono 600, color `--mint-ink` |
| `--t-body` | Body copy | 17px / 1.65 | Inter 400 |
| `--t-lead` | TL;DR / intro lead | 19px / 1.6 | Inter 400, color `--ink-soft` |
| `--t-small` | Captions, metadata | 13px / 1.5 | Inter 400, color `--ink-faint` |
| `--t-mono` | Figures, dates, versions | 14px / 1.5 | JetBrains Mono 400, `tabular-nums` |
| `--t-nav` | Sidebar nav items | 14px / 1.9 | Inter 450 |

**Measure:** body column caps at `66ch` (~640px). Never wider; long lines kill readability and
the editorial feel.

---

## 4. Spacing + layout tokens

8px base grid.

| Token | Value | Use |
|---|---|---|
| `--s-1` | 4px | hairline gaps |
| `--s-2` | 8px | tight inline |
| `--s-3` | 16px | paragraph rhythm |
| `--s-4` | 24px | between elements |
| `--s-5` | 40px | between sections |
| `--s-6` | 64px | major section breaks |
| `--s-7` | 96px | page top padding (desktop) |
| `--sidebar-w` | 232px | left nav column |
| `--content-max` | 640px | body measure (66ch) |
| `--gutter` | 56px | sidebar-to-content gap (desktop) |

**Grid:** fixed-width left sidebar (`--sidebar-w`), then a single centered content column capped
at `--content-max`. Sidebar is sticky on desktop, collapses to a top row under 880px.

| Radius | Value | Use |
|---|---|---|
| `--r-sm` | 4px | code, tags, inline marks |
| `--r-md` | 8px | callouts, cards |

Borders are hairlines: `1px solid var(--rule)`. The accent rule is `2px solid var(--mint)`.
No drop-shadows anywhere; depth comes from rules and the `--paper-2` lift, not blur. A research
letter does not float.

---

## 5. Component patterns

### Header / logo

- A left-aligned lockup: the monochrome Neko mark (filled `--ink`) + the wordmark
  **`console·log`** in Fraunces, with the dot rendered in `--mint`. Underneath, a mono kicker:
  `THE CONSOLE LABS MONTHLY` in `--t-eyebrow`.
- Sits at the top of the sidebar, not in a full-width banner. The page has no heavy header bar;
  a single hairline under the lockup is enough.

### Sidebar nav

- Sections lowercase, mono eyebrow labels (`pages`, `writing`). Items in `--t-nav`.
- Active item: `--ink` text with a 2px `--mint` left-border and `--s-3` left padding. Inactive:
  `--ink-faint`, no border. Hover lifts to `--ink-soft`.
- Nav order per spec: About, Philosophy, Purpose, experiments, blog, playbook. Group the first
  three under a `pages` label, the last three under a `writing` label.
- Sticky on desktop. Under 880px it becomes a horizontal wrap row above the content.

### Post body

- Single `--content-max` column. `--t-body`, `--ink`, line-height 1.65, paragraph spacing `--s-3`.
- Post meta line under the title: mono, `--ink-faint`. `JUN 2026 · neko-san · 6 min`. Tabular nums.
- A hairline `--rule` separates the title block from the body.

### Headings

- **h1 (post title):** `--t-display`, Fraunces, `--ink`. (The current CSS hides h1; here it is
  the centerpiece, shown.)
- **h2 (sections):** each opens with a mono eyebrow in `--mint-ink` (`TL;DR`, `SHIPPED`,
  `THINKING`), then the human heading in Fraunces below it. The eyebrow is the finance-letter tell.
- **h3:** Inter 650, `--ink`. No accent.

### Links

- Body links: `--mint-ink`, underline with `text-underline-offset: 3px` and a `--rule`-thin
  decoration that thickens to `--mint` on hover. Visited = same (a journal is evergreen).
- Nav and structural links never underline.

### Code + figures

- Inline `code`: JetBrains Mono, `--paper-2` fill, `--ink` text, `--r-sm`, `0.85em`, `2px 6px`
  padding. (Current site inverts to a black chip; we soften to a light surface so code sits
  inside prose like a research letter, not like a terminal.)
- Block `pre`: `--paper-2` surface, `--ink` text, 1px `--rule` border, `--r-md`, `--s-4` padding,
  a thin `--mint` left-border to mark it as a code block. Scrollable, `--t-mono`.
- **Figures / data callouts:** a bordered strip with a big mono number (`--t-display` size, mono,
  `tabular-nums`) and a `--t-small` caption. Use for "X swaps", "$Y backing", "Z PRs merged".

### Callouts

Three flavors, all `--r-md`, 1px border, `--s-4` padding, with a mono eyebrow label:

| Callout | Fill | Border / eyebrow | Use |
|---|---|---|---|
| **Shipped** ✅ | `--mint-wash` | `--mint` / `SHIPPED` | a release worth flagging |
| **Tinkering** 🧪 | `--amber-wash` | `--amber-ink` / `HEADS UP` | honest-about-failure, "this flopped" |
| **Note** | `--paper-2` | `--rule` / `NOTE` | neutral aside |

### Pull-quotes

- Set off from the column: Fraunces, ~24px, `--ink`, a 3px `--mint` left-border, `--mint-wash`
  fill, `--s-4` padding. The one place a sentence is allowed to be loud.
- Optional small Neko mark or `— neko-san` attribution in mono `--t-small` underneath.

### Tables

- Hairline `--rule` borders, no zebra. Header row in `--t-eyebrow` mono. Numeric columns
  `tabular-nums`, right-aligned. This is the finance default: read a table like a balance sheet.

---

## 6. How this expresses each pillar + the voice

| Pillar | Visual expression |
|---|---|
| **Small but Sharp** 🔍 | Tight measure, hairline rules, no shadows or gradients, one accent. The restraint *is* the sharpness. Nothing decorative survives that does not carry meaning. |
| **Financial-Centric** 💰 | Mono `tabular-nums` for every figure, balance-sheet tables, the data-callout strip with a big mono number, the research-letter eyebrow kickers. It reads like something a treasury desk would publish. |
| **Accessible** 🌐 | WCAG-AA throughout (every token pairing is checked), a single readable measure, system-font fallbacks, a layout that collapses cleanly to one column on mobile. Open and legible, no gatekeeping in the typography. |
| **Gamified** 🎮 | The seasoning layer: an emoji per section heading, the Neko byline, the playful `console·log` wordmark with its mint dot, the "this flopped" amber callout that turns honesty into a recurring bit. Playful in the chrome, serious in the type. |

**Voice fit (`neko-san`, first-person-plural, honest about failure):** the editorial frame gives
the casual voice somewhere credible to live. "Some stuff rocked, some didn't" lands better in a
sharp research-letter setting than in a generic blog, the contrast is the charm. The amber
"HEADS UP / this flopped" callout is the design making room for the house style of admitting what
broke. Emoji-as-punctuation maps exactly to one-emoji-per-heading.

---

## 7. Do / Don't

**Do**

- Use mint for structure: eyebrows, the active-nav border, the pre/pull-quote left-rule, inline marks.
- Set every number, date, version, and percentage in JetBrains Mono with `tabular-nums`.
- Keep the body to a 66ch measure and lean on whitespace between sections.
- Let one emoji punctuate a heading. Let Neko sign off.
- Keep tables hairline and balance-sheet-plain.
- Pick one heading family (Fraunces *or* Inter-bold) and hold it site-wide.

**Don't**

- Don't put `#45f1a6` mint on white as *text*, it fails contrast. Use `--mint-ink` for accent text.
- Don't fill large areas with mint or use it as a background band. It is a scalpel.
- Don't add drop-shadows, gradients, glows, or neon. Depth is rules + the `--paper-2` lift.
- Don't widen the body column past `--content-max`. Long lines break the editorial read.
- Don't let emoji cluster (more than one per heading reads as a chat, not a journal).
- Don't use em dashes anywhere (house rule). Commas, colons, periods.
- Don't bring the kawaii yellow Neko here. The colored cat lives on sticker.console.so; the log's
  Neko is monochrome ink.
