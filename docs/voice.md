# Console Labs voice + monthly-report skeleton

The reference the writing tool and any human author follows when drafting for
log.console.so. Distilled from `vault/purpose.md`, `vault/philosophy.md`,
`vault/about.md`.

## Who is talking

Console Labs: a small, finance-obsessed, gamer-at-heart software lab, sibling to
Dwarves Foundation. We build "small but sharp" tools for on-chain finance and we
narrate our own adventure in public. Default author persona: `neko-san`.

## The four pillars (every post should echo at least one)

- **Small but Sharp** 🔍: simple tools that nail one job.
- **Financial-Centric** 💰: safe, accessible money tooling.
- **Accessible** 🌐: open to everyone, no gatekeeping.
- **Gamified** 🎮: turn interactions into something fun.

## Voice rules

Do:
- First-person plural ("we shipped", "we learned", "we broke").
- Casual and playful: "rad", "the scoop", "vibe check", "spilling the tea". Light slang, not forced.
- Emoji as punctuation, not decoration. One per heading is plenty.
- Be honest about what failed. "Some stuff rocked, some didn't" is the house style.
- Short paragraphs. Concrete over abstract.

Don't:
- No corporate / press-release register ("We are pleased to announce").
- No hype without substance. No empty superlatives.
- No em dashes (use commas, colons, or a sentence break).
- Don't leak unannounced strategy, custodial-key details, infra internals, or holder-comms that are still embargoed (same privacy gate as the rest of Console Labs).

## Monthly report skeleton

Front matter (Hugo, via the `New Post` templater):

```yaml
---
tags: [consolelabs, work, monthly]
title: "Console Log: <Month Year>"
date: <YYYY-MM-DD>
description: <one-line hook in the voice>
authors: [neko-san]
toc: false
type:
---
```

Body sections (drop any that have nothing real to say; never pad):

1. **TL;DR** : 2-3 sentences, what mattered this month.
2. **Shipped** : releases + notable changes, each with a one-line "why it matters".
3. **Tinkering** : experiments, spikes, things we tried (include the ones that flopped).
4. **Thinking** : the thought-process / decisions / shifts in direction. This is the self-narration heart.
5. **Next** : what we are chasing next month. Light, not a roadmap commitment.

Keep it skimmable. A monthly log is a story, not a changelog dump; the changelog
is the evidence, the narration is the point.

## Source material for a given month

A monthly report is assembled from real signal, not invented:
- merged PRs / releases across the active Console Labs repos that month,
- experiment + research notes,
- decisions recorded in the tenant repos.

The writing tool's job is to gather that signal, then narrate it in this voice.
