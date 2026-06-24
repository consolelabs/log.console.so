# Brand Direction: Arcade Neko

A token-based visual direction for **log.console.so**, the monthly lab journal of Console Labs. Author persona: `neko-san`.

The angle is in the name. Console Labs runs a **console**. The existing logo is already a cat's head whose face is a gamepad. Arcade Neko leans all the way into that: a retro arcade / terminal aesthetic, a monospace voice, a punchy palette, and the black Neko cat as a real character that lives in the margins of the page. It is fun and loud where the current site is silent, but it never trades away long-form readability. A monthly log is a story you read top to bottom; the theme is the set dressing, not the obstacle.

This is a reaction against the current styling (white page, thin sans, one mint accent, very quiet). We keep the mint, keep the black cat, keep the minimal bones, then give them a CRT and a coin slot.

---

## 1. Design principles

1. **The page is a console screen.** Dark surface, glowing text, a thin scanline texture, a chunky bezel around the reading column. Content sits "inside" the machine.
2. **Mono for the chrome, humanist sans for the prose.** Anything that is interface (nav, labels, metadata, code, callout headers, the logo wordmark) is monospace. Anything you read for more than ten seconds (post body) is a readable sans. This split is the whole legibility strategy.
3. **One hot accent, used like a power LED.** Mint-green stays the brand signal, but it is now a glow, not a flat fill. Reserve it for active states, headings, the cursor, and the cat's eyes. Scarcity keeps it loud.
4. **The cat has agency.** Neko is not a static logo. It blinks in the header, sits at the bottom of the sidebar, and shows up as a `>` prompt before section titles. It is a character, not a watermark.
5. **8-bit as seasoning, not the meal.** Pixel corners, a coin-slot divider, a press-start button, a scanline. Used sparingly so the body text never inherits the pixelation.

---

## 2. Color tokens

Dark-first (the default). A light "day mode" inverts the surfaces and keeps the same accent.

| Token | Hex | Role |
|---|---|---|
| `--c-bg` | `#0d1117` | Page background (deep console charcoal, not pure black, easier on eyes for long reads) |
| `--c-bg-raised` | `#161b22` | Sidebar, callout, code-block, post "cartridge" surface |
| `--c-bezel` | `#21262d` | Borders, the bezel frame, dividers, hairlines |
| `--c-text` | `#e6edf3` | Body text (contrast 14.0:1 on `--c-bg`, well past AA) |
| `--c-text-dim` | `#9da7b3` | Metadata, captions, secondary nav (contrast 6.9:1, passes AA) |
| `--c-mint` | `#45f1a6` | Primary accent. Headings, links, cursor, cat eyes, active nav (contrast 12.6:1 on `--c-bg`) |
| `--c-mint-deep` | `#2ea370` | Mint for large fills / hover where the bright mint would vibrate |
| `--c-coin` | `#ffcb47` | Secondary "arcade gold". Coin slot, the Gamified pillar, emoji-tier highlights. Use rarely (contrast 11.0:1) |
| `--c-berry` | `#ff5d8f` | Tertiary alert pink. "we broke it" / failure callouts, honest-about-failure moments (contrast 7.1:1) |
| `--c-link-visited` | `#7ee3bf` | Visited links, a desaturated mint |

**Contrast contract (WCAG-AA, the non-negotiable):**
- Body `--c-text` on `--c-bg`: 14.0:1 ✓ (AAA)
- Body on `--c-bg-raised`: 12.3:1 ✓ (AAA)
- Dim text `--c-text-dim` on `--c-bg`: 6.9:1 ✓ (AA, and AAA for large)
- Links `--c-mint` on both surfaces: > 11:1 ✓
- Never put `--c-coin` or `--c-mint` text under 1.2rem on a raised surface without re-checking; the bright accents are for headings, labels, and chrome, not body runs.

**Light "day mode" (the toggle, optional but specced):**
| Token | Hex |
|---|---|
| `--c-bg` | `#fafdfb` |
| `--c-bg-raised` | `#f0f4f1` |
| `--c-bezel` | `#d4ddd6` |
| `--c-text` | `#11221a` |
| `--c-text-dim` | `#3d4f47` |
| `--c-mint` | `#1f8f5f` (the bright mint fails AA on white, so day mode uses the deeper mint for text; 4.9:1 ✓) |

---

## 3. Type tokens

Two families. A monospace for all interface + structure, a clean sans for reading.

| Token | Stack | Use |
|---|---|---|
| `--font-mono` | `"JetBrains Mono", "Space Mono", ui-monospace, "SF Mono", Menlo, monospace` | Logo wordmark, nav, headings, labels, metadata, code, callout headers, buttons |
| `--font-body` | `"Inter", "IBM Plex Sans", system-ui, -apple-system, sans-serif` | Post body, long-form paragraphs, lists |
| `--font-display` | `"Press Start 2P", "JetBrains Mono", monospace` | The 8-bit accent only: the masthead tagline, the "PRESS START" button, score-style numerals. Never for a run of text longer than a few words (it is illegible at length, that is the point of restricting it) |

Google Fonts link (mockup): JetBrains Mono (400/500/700), Inter (400/500/600), Press Start 2P (400).

**Scale** (1.25 major-third-ish, tuned for a mono + sans pair):

| Token | Size / line-height | Element |
|---|---|---|
| `--t-display` | `clamp(1.6rem, 4vw, 2.4rem)` / 1.2 | Masthead `console_log` wordmark |
| `--t-h1` | `1.85rem` / 1.25 | Post title |
| `--t-h2` | `1.3rem` / 1.3 | Section headers (TL;DR, Shipped, Thinking) |
| `--t-h3` | `1.05rem` / 1.35 | Sub-headers |
| `--t-body` | `1.0rem` (16px) / **1.7** | Post body. Generous leading is the readability insurance against the dark theme |
| `--t-small` | `0.82rem` / 1.5 | Metadata, captions, nav items |
| `--t-code` | `0.85rem` / 1.55 | Inline + block code |
| `--t-tag` | `0.72rem` / 1 | Uppercase mono labels (`// SHIPPED`, tag pills) |

Body measure caps at **66ch** (`--measure: 66ch`). This is the single most important readability lever on a dark long-form page.

Heading treatment: section headers render lowercase mono with a mint `> ` prompt prefix and a `_` blinking-cursor option, e.g. `> tl;dr_`. Post title is sentence case, mono, mint.

---

## 4. Spacing + radius + effect tokens

Spacing on a 4px base (an "8-bit" grid).

| Token | Value |
|---|---|
| `--sp-1` | `4px` |
| `--sp-2` | `8px` |
| `--sp-3` | `16px` |
| `--sp-4` | `24px` |
| `--sp-5` | `40px` |
| `--sp-6` | `64px` |
| `--main-padding` | `40px` (desktop) / `20px` (mobile) |
| `--nav-width` | `220px` |
| `--measure` | `66ch` |

| Token | Value | Note |
|---|---|---|
| `--radius-chip` | `2px` | Pixel-sharp corners on chrome (pills, buttons, code) |
| `--radius-card` | `6px` | Slightly softer on the big "cartridge" cards |
| `--border-hair` | `1px solid var(--c-bezel)` | Default hairline |
| `--border-bezel` | `2px solid var(--c-bezel)` | The console-frame border |
| `--glow-mint` | `0 0 0 1px var(--c-mint), 0 0 12px -2px var(--c-mint)` | Focus / active ring, used sparingly |
| `--scanline` | `repeating-linear-gradient(0deg, transparent 0 3px, rgba(255,255,255,.012) 3px 4px)` | CRT texture overlay, ultra-faint so it never hurts text |

---

## 5. Component patterns

### Header / logo
- Full-width top bar on `--c-bg`, a `--border-bezel` bottom edge.
- **Logo**: the existing cat-with-gamepad-face mark, recolored. Cat outline + ears in `--c-text`, the gamepad face filled, the two D-pad/eye dots in `--c-mint` (the cat's "eyes" glow). Beside it, a mono wordmark `console_log` with a blinking `_` cursor.
- A small mono breadcrumb on the right: `~/labs/log $` to reinforce the terminal frame.
- The cat blinks: the mint eyes drop opacity every ~5s (CSS keyframe). Optional, degrade to static.

### Nav (left sidebar)
- Mono, lowercase, `--t-small`. Sections grouped: the "lab" group (about, philosophy, purpose) and the "stream" group (experiments, blog, playbook).
- Each item is prefixed with a dim `> `. Active item: mint text + a 2px left mint bar + the prompt turns bright.
- Hover: text goes mint, a faint `--glow-mint`-lite underline animates in.
- Bottom of the sidebar: a small seated Neko cat ASCII/SVG mascot with a one-line mono caption like `// neko-san is watching the build`.
- Above the cat: an arcade "score" stat block, e.g. `SHIPPED: 04  BROKE: 01  XP: +120`. Pure flavor, expresses Gamified.

### Post body
- Wrapped in a "cartridge" card: `--c-bg-raised`, `--radius-card`, `--border-bezel`, with a faint top notch (a pixel coin-slot) to read as a game cartridge / console screen.
- Body text: `--font-body`, `--t-body`, line-height 1.7, max-width `--measure`. This is the legibility core; the cartridge is just the frame.
- First section (TL;DR) gets a subtle mint left rule to anchor the eye.

### Headings
- `h1` post title: mono, mint, `--t-h1`.
- `h2` section header: lowercase mono with `> ` mint prompt and a trailing `_`. A coin-slot divider (`■ ■ ■` dotted mono rule) sits under each, gold-tinted, very thin.
- Each section header carries its pillar emoji inline (🔍 💰 🌐 🎮), one max, per the voice rules.

### Links
- Body links: `--c-mint`, no underline at rest, a mint underline (1px) animates in on hover. Visited: `--c-link-visited`.
- Focus-visible: `--glow-mint` ring (keyboard accessibility, not optional).
- External links get a tiny mono `↗`.

### Code + terminal callout
- Inline code: `--font-mono`, `--c-mint` text on `--c-bg-raised`, `--radius-chip`, 2px/6px padding.
- Block code / terminal: a "console window" with a 3-dot title bar (the dots in mint/coin/berry, an arcade nod), `--c-bg` (darker than the cartridge so it recedes), mono, a mint `$` prompt on command lines, dim text for output. A faint `--scanline` overlay only here.

### Callouts
- **Note** (default): `--c-bg-raised`, mint left bar, a `// note` mono label.
- **Win / shipped**: mint-tinted, label `// shipped 🚀`.
- **Broke it** (honest-about-failure, a house value): `--c-berry` left bar + `// we broke it 💥` label. Failure gets its own loud color on purpose; it is part of the brand voice.
- **Coin / fun fact**: `--c-coin` accent, `// bonus 🪙`.

### Buttons / pills
- Tag pills: `--t-tag`, uppercase mono, `--c-bg-raised`, `--c-bezel` border, `--radius-chip`.
- Primary CTA ("read the full log", subscribe): a `--font-display` "PRESS START"-style button, mint border + mint text, fills mint with dark text on hover, a 2px pixel-offset shadow that snaps on press (`:active` translates 2px). The one place Press Start 2P is allowed.

---

## 6. How this expresses the four pillars + the voice

| Pillar | Visual expression |
|---|---|
| **Small but Sharp** 🔍 | Tight 66ch measure, one accent, monospace precision, pixel-sharp 2px corners. Nothing decorative survives that does not earn its place. The restraint *is* the sharpness. |
| **Financial-Centric** 💰 | The arcade "score" block reads like a ledger (`SHIPPED / BROKE / XP`). Tabular mono numerals everywhere figures appear. Treasury / shipped counts get the gold `--c-coin` so money-shaped data is visually distinct. |
| **Accessible** 🌐 | AA-and-mostly-AAA contrast is a hard token contract, not an afterthought. A day-mode toggle. Sans body at 1.7 line-height and a capped measure keep long reads easy. The theme never wins over the reader. |
| **Gamified** 🎮 | The whole frame: console bezel, coin-slot dividers, the PRESS START button, the score block, the blinking cat, scanlines, the `>` prompt. The literal Console = console pun, made the centerpiece. |

**Voice fit** (playful, casual, emoji-friendly, honest about failure, first-person-plural): the terminal prompt voice (`~/labs/log $`) matches "we shipped / we broke". The dedicated berry "we broke it" callout bakes the honest-about-failure value straight into the component set. One emoji per heading honors the "emoji as punctuation" rule. The cat-as-character carries the casual, fun register without a single word of copy.

---

## 7. Do / Don't

**Do**
- Keep body text in `--font-body` sans at 1.7 line-height, capped at 66ch. Readability is the brand's first promise.
- Reserve mint for accent + active states. Treat it like a power LED, not a paint bucket.
- Use the cat as a recurring character (header blink, sidebar sit, `>` prompt), not a static stamp.
- Let failure be loud: use the berry "we broke it" callout honestly and often. It is on-voice.
- Keep `--c-coin` gold for money-shaped + score-shaped data only.
- Verify every accent-on-surface pair against AA before shipping a new component.

**Don't**
- Don't set body paragraphs in monospace or in Press Start 2P. Mono is for chrome; pixel font is for ≤ a few words.
- Don't tint body text mint or coin. Accents are for headings, labels, links, and chrome.
- Don't stack effects: scanlines + glow + pixel-shadow on the same element is noise. One effect per surface.
- Don't go pure-black `#000`; `--c-bg` is `#0d1117` so the glow and the long read both behave.
- Don't add a second hot accent beyond mint/coin/berry. Three accents is the ceiling.
- Don't let the arcade chrome shrink the reading column below 60ch on desktop.
- No em dashes anywhere (house rule). Commas, colons, or a sentence break.
