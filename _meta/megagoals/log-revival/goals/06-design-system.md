# Sub-goal 06: design system + brand guideline (DESIGN.md)

**Merge policy:** auto
**Time budget:** 1-2 calendar days of loop work (parallel designer sub-agents + critique rounds)
**Proof:** `docs/brand/DESIGN.md` (locked) + the full design trail under `docs/brand/`: current-state captures, 2-3 explored direction boards, the critique/selection rationale, and a sample log-homepage mockup rendered in the chosen direction
**Depends on:** none
**Branch:** docs/log-revival-06-design
**PR base:** main

## Outcome

A `docs/brand/DESIGN.md` brand guideline for Console Labs' web surfaces, produced by a rigorous multi-step design process (not a one-shot), and validated on a real mockup. It locks color tokens, a type scale, spacing, component patterns, logo/mark usage, and a voice-to-visual mapping. It is the single source of truth the log reskin (07) implements. Because there is no human gate mid-process, the rigor IS the safeguard: diverge, critique, converge, validate.

## Quality bar

A designer could hand this to any engineer and get a consistent, on-brand surface. Opinionated and token-based, not a mood board. It feels like Console Labs (playful, sharp, finance-native) and stays coherent with the existing sticker brand, not a generic template. The chosen direction is defensible: the trail shows what else was considered and why this won.

## How to close the loop (multi-step; do not collapse into one shot)

1. **Audit current state.** Capture the live styling of sticker.console.so + log.console.so: screenshots plus the actual tokens (colors, fonts, spacing) pulled from their CSS / Hugo layouts. Locate sticker's source (its repo if one exists, else the live site). Write `docs/brand/current-state.md` and store captures under `docs/brand/visuals/current/`.
2. **Diverge.** Dispatch 2-3 UI/UX designer sub-agents IN PARALLEL (the `frontend-design` + `website-brand` skills), each proposing a DISTINCT on-brand direction grounded in `docs/voice.md` (the four pillars) + the current-state audit. Each emits a direction board (palette + type + component vibe + a one-screen sketch) under `docs/brand/directions/<name>/`.
3. **Critique + select.** Run the kit's `/visual-team` lenses (or dispatch critic sub-agents) across the directions on distinct axes (brand-fit, voice-match, distinctiveness, accessibility). Score each; pick the strongest, grafting the best elements of the runners-up. Record the decision in `docs/brand/RATIONALE.md` (this is your call; justify it).
4. **Converge + lock.** Write `docs/brand/DESIGN.md` from the winner: color tokens, type scale, spacing scale, component patterns, logo/mark do + don't, and the voice-to-visual section.
5. **Validate.** Produce a sample log-homepage mockup in the locked direction (static HTML/CSS or a rendered preview) to prove the system works on real content and to give sub-goal 07 a concrete target. Screenshot it into `docs/brand/visuals/mockup/`.

**Done =** `docs/brand/DESIGN.md` exists with LOCKED color + type + spacing + component tokens and a voice-to-visual section, AND the full trail is committed: `current-state.md`, 2-3 `directions/`, `RATIONALE.md`, and the sample homepage mockup.

## Scope edges

**In:** the multi-step design process, `docs/brand/DESIGN.md`, the design trail (current-state, directions, rationale, mockup).
**Out:** implementing the reskin on the real layouts (that is 07), the writing tool, content, the build pipeline.
**Not:** redesigning the sticker site, building a full component library / Storybook, inventing a new logo from scratch (work from the existing mark), rebranding Console Labs' identity wholesale, touching the live log layouts yet, collapsing the diverge/critique steps into a single generated guideline.

## Where to look

`layouts/` + `assets/` (current log styling), sticker.console.so (live site or its repo), `docs/voice.md` (the voice to express visually), the `frontend-design` + `website-brand` skills and `/visual-team`.

## PR body

A Console Labs web brand guideline at `docs/brand/DESIGN.md`, produced via a multi-step design process (audit, parallel divergent directions from UI/UX designer sub-agents, /visual-team critique, convergence, mockup validation). Source of truth for the log reskin (07).

Verify: `docs/brand/DESIGN.md` carries locked tokens + a voice-to-visual section; the full trail (current-state, directions, rationale, mockup) is under `docs/brand/`.

Part of mega-goal `log-revival`.

## Notes
