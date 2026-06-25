# Brand-direction selection rationale

Three directions were explored in parallel (each by a separate designer pass),
grounded in `docs/voice.md` (the four pillars) and the `current-state.md` audit.
This file records the scoring and the call. The winner is **fintech-editorial**,
grafting elements from both runners-up.

## The three directions

| Direction | One-line angle |
|---|---|
| **mint-sharpened** | The current mint-and-white minimal, made deliberate: a real type scale + spacing grid, accent demoted to structural-only. An evolution. |
| **fintech-editorial** | A sharp crypto-research letter with a sense of humor. Serif headings, mono figures, balance-sheet tables, one mint scalpel. Finance literacy in the details. |
| **arcade-neko** | A retro CRT / arcade console: dark surface, scanlines, mono chrome, the Neko as a live character with a `>` prompt. Bold and loud. |

## Scoring (1-5, higher is better)

Axes from the 06 contract: brand-fit (coherent with sticker + the mint DNA),
voice-match (expresses the four pillars + neko-san), distinctiveness (its own
identity, not a template), accessibility (WCAG-AA, focus states), plus
long-form-readability (a monthly journal is read top to bottom).

| Axis | mint-sharpened | fintech-editorial | arcade-neko |
|---|---|---|---|
| Brand-fit | 5 | 4 | 3 |
| Voice-match | 3 | 5 | 4 |
| Distinctiveness | 2 | 4 | 5 |
| Accessibility | 5 | 5 | 4 |
| Long-form readability | 5 | 5 | 3 |
| **Total** | **20** | **23** | **19** |

## The call: arcade-neko (Han's pick)

The scoring put fintech-editorial highest on paper (23) and my loop recommendation
was fintech-editorial. **Han overrode it and picked arcade-neko** (2026-06-25). This
is his brand to own, and the override is decisive: arcade-neko is the locked
direction. The trade he is accepting (boldness + the literal Console = console pun
over the safest readability play) is a deliberate identity choice, and arcade-neko
is the only direction that turns the lab's actual identity (a console, a gamepad
cat) into the design language instead of treating it as decoration.

It is also the most **distinctive** (top score, 5/5): nobody mistakes it for a
template or a generic dev blog. It expresses all four pillars hard, the score block
as a ledger (Financial-Centric), the console frame + coin slots + PRESS START
(Gamified), the 66ch sans body at 1.7 leading (Accessible), one mint accent used
like a power LED (Small but Sharp). The readability risk that dropped it on my
scorecard is mitigated by the direction's own hard rules: body stays a clean sans,
never mono or pixel font; the measure is capped; the bg is `#0d1117` not pure black;
one effect per surface. Those rules are now load-bearing in DESIGN.md.

**Why not fintech-editorial:** higher on the paper rubric and the safer long-form
read, but it is the more generic identity (a research letter is a known shape) and
it under-uses the Console pun that is sitting right there in the logo. Han chose
identity over safety.

**Why not mint-sharpened:** safest and most coherent, lowest on distinctiveness and
voice. It sharpens the current site without giving the lab an identity of its own.

## Grafts from the runners-up

The locked direction (arcade-neko) absorbs the strongest non-conflicting ideas from
the others, applied as guardrails against the dark-theme readability risk:

- **From fintech-editorial:** the discipline that keeps the arcade frame from
  hurting the read, the hard 66ch measure, mono `tabular-nums` on every figure,
  hairline balance-sheet tables, and "the body is sacred, the chrome is where the
  fun lives." DESIGN.md keeps the body a clean sans at 1.7 leading for exactly this.
- **From mint-sharpened:** accent strictly structural (mint as a power LED, not a
  paint bucket), hierarchy through type + space, every gap a step on the grid.

The locked spec is `docs/brand/DESIGN.md`.
