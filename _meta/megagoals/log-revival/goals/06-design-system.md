# Sub-goal 06: design system + brand guideline (DESIGN.md)

**Merge policy:** gate
**Time budget:** 1-2 calendar days, mostly designer-sub-agent iteration + Han's taste approval
**Proof:** `docs/brand/DESIGN.md` committed, plus 2-3 visuals: current-state captures of sticker.console.so + log.console.so, and a proposed palette + type board
**Depends on:** none
**Branch:** docs/log-revival-06-design
**PR base:** main

## Outcome

A `docs/brand/DESIGN.md` brand guideline for Console Labs' web surfaces, synthesized from the CURRENT styling of sticker.console.so and log.console.so and sharpened by a UI/UX designer sub-agent. It locks color tokens, a type scale, spacing, component patterns, logo/mark usage, and a voice-to-visual mapping (how small-but-sharp / gamified / finance-native reads visually). It is the single source of truth the log reskin (07) implements.

## Quality bar

A designer could hand this to any engineer and get a consistent, on-brand surface. Opinionated and token-based, not a mood board. It feels like Console Labs (playful, sharp, finance-native) and stays coherent with the existing sticker brand, not a generic template. Han recognizes it as the lab's identity, leveled up.

## How to close the loop

- Capture current state: pull the live styling from sticker.console.so and log.console.so (screenshots + the actual CSS / Hugo layout tokens). Locate sticker's source (its repo if one exists, else the live site). Store reference captures under `docs/brand/visuals/`.
- Dispatch a UI/UX designer sub-agent (the `frontend-design` skill, and the `website-brand` skill for the brand-token workflow) to synthesize a direction from those captures + the voice in `docs/voice.md`. Critique it through the kit's `/visual-team` lenses and iterate to a locked token set.
- Write `docs/brand/DESIGN.md`: color tokens, type scale, spacing scale, component patterns, logo/mark do + don't, and the voice-to-visual section.

**Done =** `docs/brand/DESIGN.md` exists with LOCKED color + type + spacing + component tokens and a voice-to-visual section, AND the current-state captures + a proposed palette/type board are committed under `docs/brand/visuals/`.

## Scope edges

**In:** `docs/brand/DESIGN.md`, brand-token synthesis, the current-state reference captures, the designer sub-agent + visual-team critique.
**Out:** implementing the reskin (that is 07), the writing tool, content, the build pipeline.
**Not:** redesigning the sticker site, building a full component library / Storybook, inventing a new logo from scratch (work from the existing mark), rebranding Console Labs' identity wholesale, touching the live log layouts yet.

## Where to look

`layouts/` + `assets/` (current log styling), sticker.console.so (live site or its repo for current sticker styling), `docs/voice.md` (the voice to express visually), the `frontend-design` + `website-brand` skills and `/visual-team`.

## PR body

A Console Labs web brand guideline at `docs/brand/DESIGN.md`, synthesized from the current sticker.console.so + log.console.so styling via a UI/UX designer sub-agent and locked to token form. Source of truth for the log reskin (07).

Verify: `docs/brand/DESIGN.md` carries locked color/type/spacing/component tokens + a voice-to-visual section; reference captures + palette board under `docs/brand/visuals/`.

GATE: Han approves the brand direction. Part of mega-goal `log-revival`.

## Notes
