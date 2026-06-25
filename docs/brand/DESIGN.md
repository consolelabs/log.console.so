# Console Labs web design system (DESIGN.md)

**Status: LOCKED.** This is the single source of truth for log.console.so's visual
brand. Sub-goal 07 implements these tokens in the Hugo `layouts/` + `assets/`. The
selection trail (audit, three directions, scoring) is in `current-state.md`,
`directions/`, and `RATIONALE.md`. Do not re-open the brand decisions here; refine
via a new design pass if needed.

Direction: **Arcade Neko.** Console Labs runs a *console*, and the Neko mark is
already a cat whose face is a gamepad. This leans all the way in: a retro arcade /
terminal aesthetic, a monospace voice for all chrome, a dark "console screen"
surface, and the black Neko as a recurring character that lives in the page
margins. It is loud and fun where the current site is silent, but readability is
non-negotiable: the body stays a clean sans at a capped measure. The theme is set
dressing, never the obstacle. We keep the mint, keep the black cat, keep the
minimal bones, then give them a CRT and a coin slot.

## 1. Color tokens

Dark-first (the default). A light "day mode" inverts the surfaces, same accent.

### Dark mode (default)

| Token | Hex | Role |
|---|---|---|
| `--c-bg` | `#0d1117` | Page background (deep console charcoal, NOT pure black) |
| `--c-bg-raised` | `#161b22` | Sidebar, callout, code-block, post "cartridge" surface |
| `--c-bezel` | `#21262d` | Borders, the bezel frame, dividers, hairlines |
| `--c-text` | `#e6edf3` | Body text (14.0:1 on `--c-bg`, AAA) |
| `--c-text-dim` | `#9da7b3` | Metadata, captions, secondary nav (6.9:1, AA) |
| `--c-mint` | `#45f1a6` | Primary accent: headings, links, cursor, cat eyes, active nav (12.6:1) |
| `--c-mint-deep` | `#2ea370` | Mint for large fills / hover where bright mint would vibrate |
| `--c-coin` | `#ffcb47` | "Arcade gold": coin slot, money + score data, the Gamified pillar. Rare (11.0:1) |
| `--c-berry` | `#ff5d8f` | Alert pink: "we broke it" / honest-about-failure callouts (7.1:1) |
| `--c-link-visited` | `#7ee3bf` | Visited links (desaturated mint) |

### Light "day mode" (toggle)

| Token | Hex |
|---|---|
| `--c-bg` | `#fafdfb` |
| `--c-bg-raised` | `#f0f4f1` |
| `--c-bezel` | `#d4ddd6` |
| `--c-text` | `#11221a` |
| `--c-text-dim` | `#3d4f47` |
| `--c-mint` | `#1f8f5f` (bright mint fails AA on white; day mode uses the deeper mint for text, 4.9:1) |

**Contrast contract (WCAG-AA, non-negotiable):** body on `--c-bg` 14.0:1, body on
`--c-bg-raised` 12.3:1, dim text 6.9:1, mint links >11:1. Never put `--c-coin` or
`--c-mint` text under 1.2rem on a raised surface without re-checking; bright accents
are for headings, labels, and chrome, never body runs.

## 2. Type tokens

Two families plus one 8-bit accent. Mono for all interface + structure, a clean
sans for reading.

| Token | Stack | Use |
|---|---|---|
| `--font-mono` | `"JetBrains Mono", "Space Mono", ui-monospace, "SF Mono", Menlo, monospace` | Logo wordmark, nav, headings, labels, metadata, code, callout headers, buttons |
| `--font-body` | `"Inter", "IBM Plex Sans", system-ui, -apple-system, sans-serif` | Post body, long-form paragraphs, lists |
| `--font-display` | `"Press Start 2P", "JetBrains Mono", monospace` | 8-bit accent ONLY: masthead tagline, "PRESS START" button, score numerals. Never a run longer than a few words |

Google Fonts (mockup): JetBrains Mono 400/500/700, Inter 400/500/600, Press Start 2P 400.

### Scale (1.25-ish, tuned for mono + sans)

| Token | Size / line-height | Element |
|---|---|---|
| `--t-display` | `clamp(1.6rem, 4vw, 2.4rem)` / 1.2 | Masthead `console_log` wordmark |
| `--t-h1` | `1.85rem` / 1.25 | Post title (mono, mint) |
| `--t-h2` | `1.3rem` / 1.3 | Section headers (tl;dr, shipped, thinking) |
| `--t-h3` | `1.05rem` / 1.35 | Sub-headers |
| `--t-body` | `1.0rem` (16px) / **1.7** | Post body. Generous leading is the readability insurance on dark |
| `--t-small` | `0.82rem` / 1.5 | Metadata, captions, nav items |
| `--t-code` | `0.85rem` / 1.55 | Inline + block code |
| `--t-tag` | `0.72rem` / 1 | Uppercase mono labels (`// SHIPPED`, tag pills) |

**Measure:** body caps at `--measure: 66ch`. The single most important readability
lever on a dark long-form page; never below 60ch on desktop.

**Heading treatment:** section headers render lowercase mono with a mint `> `
prompt prefix and an optional `_` blinking cursor, e.g. `> tl;dr_`. Each carries
its pillar emoji inline (🔍 💰 🌐 🎮), one max. Post title is sentence-case mono mint.

## 3. Spacing + radius + effect tokens

4px base ("8-bit" grid).

| Token | Value |
|---|---|
| `--sp-1` | `4px` |
| `--sp-2` | `8px` |
| `--sp-3` | `16px` |
| `--sp-4` | `24px` |
| `--sp-5` | `40px` |
| `--sp-6` | `64px` |
| `--main-padding` | `40px` desktop / `20px` mobile |
| `--nav-width` | `220px` |
| `--measure` | `66ch` |

| Token | Value | Note |
|---|---|---|
| `--radius-chip` | `2px` | Pixel-sharp corners on chrome (pills, buttons, code) |
| `--radius-card` | `6px` | Softer on the big "cartridge" cards |
| `--border-hair` | `1px solid var(--c-bezel)` | Default hairline |
| `--border-bezel` | `2px solid var(--c-bezel)` | The console-frame border |
| `--glow-mint` | `0 0 0 1px var(--c-mint), 0 0 12px -2px var(--c-mint)` | Focus / active ring, sparingly |
| `--scanline` | `repeating-linear-gradient(0deg, transparent 0 3px, rgba(255,255,255,.012) 3px 4px)` | CRT texture, ultra-faint so it never hurts text |

**One effect per surface.** Never stack scanline + glow + pixel-shadow on the same
element. Never pure-black `#000`.

## 4. Component patterns

- **Header / logo:** full-width top bar on `--c-bg`, `--border-bezel` bottom edge.
  The cat-with-gamepad mark recolored: outline + ears in `--c-text`, gamepad face
  filled, the two D-pad/eye dots in `--c-mint` (eyes glow). Mono wordmark
  `console_log` with a blinking `_`. A mono breadcrumb on the right: `~/labs/log $`.
  The cat blinks (mint eyes drop opacity ~5s); degrade to static.
- **Nav (left sidebar):** mono lowercase `--t-small`. Two groups: "lab" (about,
  philosophy, purpose) and "stream" (experiments, blog, playbook). Each item
  prefixed with a dim `> `. Active = mint text + 2px left mint bar + bright prompt.
  Hover = text mint + faint underline. Bottom: a small seated Neko mascot with a
  mono caption (`// neko-san is watching the build`). Above it, an arcade "score"
  block: `SHIPPED: 04  BROKE: 01  XP: +120` (flavor, expresses Gamified).
- **Post body:** wrapped in a "cartridge" card: `--c-bg-raised`, `--radius-card`,
  `--border-bezel`, a faint top notch (pixel coin-slot). Body = `--font-body`,
  `--t-body`, line-height 1.7, max-width `--measure`. The cartridge is just the
  frame; the sans body is the legibility core. TL;DR gets a mint left rule.
- **Headings:** h1 (post title) mono mint `--t-h1`. h2 lowercase mono with `> `
  mint prompt + trailing `_`, a thin gold coin-slot divider (`■ ■ ■` dotted rule)
  under each. One pillar emoji per header.
- **Links:** body links `--c-mint`, no underline at rest, 1px mint underline animates
  in on hover. Visited `--c-link-visited`. Focus-visible `--glow-mint` ring (NOT
  optional, keyboard a11y). External links get a tiny mono `↗`.
- **Code + terminal:** inline code mono, `--c-mint` text on `--c-bg-raised`,
  `--radius-chip`. Block code = a "console window" with a 3-dot title bar (dots in
  mint/coin/berry), `--c-bg` (darker than the cartridge so it recedes), a mint `$`
  prompt on command lines, dim output. Faint `--scanline` overlay only here.
- **Callouts:** Note = `--c-bg-raised` + mint left bar + `// note`. Win/shipped =
  mint-tinted + `// shipped 🚀`. **Broke it** = `--c-berry` left bar + `// we broke
  it 💥` (failure gets its own loud color on purpose, it is part of the voice).
  Coin/fun-fact = `--c-coin` + `// bonus 🪙`.
- **Buttons / pills:** tag pills = `--t-tag` uppercase mono, `--c-bg-raised`,
  `--c-bezel` border, `--radius-chip`. Primary CTA = a `--font-display` "PRESS
  START" button, mint border + text, fills mint with dark text on hover, a 2px
  pixel-offset shadow that snaps on `:active`. The one place Press Start 2P is allowed.

## 5. Logo / mark usage

- The log uses the **existing cat-with-gamepad mark, recolored to ink + mint eyes**
  (outline `--c-text`, eye dots `--c-mint`). Work from the mark in `assets/img/`;
  do not invent a new logo.
- The **colored (yellow) kawaii Neko stays on sticker.console.so.** Never here. The
  log's cat is the lab's monochrome character: it blinks in the header, sits in the
  sidebar, prompts sections with `>`.
- The wordmark is `console_log` (mono, blinking `_` cursor). Do not stretch,
  recolor beyond ink + mint eyes, or stack effects on it.

## 6. Voice-to-visual mapping

| Pillar | Visual expression |
|---|---|
| **Small but Sharp** 🔍 | Tight 66ch measure, one accent, monospace precision, pixel-sharp 2px corners. The restraint is the sharpness. |
| **Financial-Centric** 💰 | The "score" block reads like a ledger (`SHIPPED / BROKE / XP`); tabular mono numerals on every figure; money + score data gets the gold `--c-coin` so it is visually distinct. |
| **Accessible** 🌐 | AA-and-mostly-AAA contrast as a hard token contract, a day-mode toggle, sans body at 1.7 leading, a capped measure. The theme never wins over the reader. |
| **Gamified** 🎮 | The whole frame: console bezel, coin-slot dividers, the PRESS START button, the score block, the blinking cat, faint scanlines, the `>` prompt. The literal Console = console pun is the centerpiece. |

**Voice fit (`neko-san`, playful, honest about failure, first-person-plural):** the
terminal prompt voice (`~/labs/log $`) matches "we shipped / we broke". The berry
"we broke it 💥" callout bakes the honest-about-failure value straight into the
component set. One emoji per heading honors emoji-as-punctuation. The cat-as-character
carries the casual register with no copy.

## 7. Do / Don't

**Do:** keep body in `--font-body` sans at 1.7 leading, capped at 66ch (readability
is the first promise); reserve mint for accent + active states (a power LED, not a
paint bucket); use the cat as a recurring character (header blink, sidebar sit, `>`
prompt); let failure be loud with the berry "we broke it" callout, honestly and
often; keep `--c-coin` gold for money + score data only; verify every accent-on-
surface pair against AA before shipping a component.

**Don't:** set body paragraphs in mono or Press Start 2P (mono = chrome, pixel font
= a few words max); tint body text mint or coin; stack effects (one per surface);
use pure-black `#000` (bg is `#0d1117`); add a fourth hot accent beyond mint/coin/
berry; shrink the reading column below 60ch on desktop; use em dashes (house rule,
commas/colons/periods).
