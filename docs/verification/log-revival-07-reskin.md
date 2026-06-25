# Verification: log-revival-07-reskin

Behavioral/visual change: reskin log.console.so's Hugo `layouts/` + `assets/` to
the `docs/brand/DESIGN.md` (Arcade Neko) brand. Proof that the site renders on the
locked tokens with zero layout overflow, plus a negative control proving the new
stylesheet is what produces the brand. Companion detail + the full token checklist:
`docs/proof/07-reskin.md`.

## Green run

```
Command: env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
Exit:    0   (content regenerated)
Verdict: PASS

Command: env -u GOROOT hugo --minify
Exit:    0   (built, 0 ERROR lines; only pre-existing languageCode-deprecation +
              content-bak duplicate-menu WARNINGS, not from this change)
Verdict: PASS

Measure: Playwright getBoundingClientRect at a 375px mobile viewport
         documentElement.scrollWidth == clientWidth == 375
         elements past viewport: 0 ; main column width: 335
Verdict: PASS  (no horizontal scroll, nothing clipped; see docs/proof/07-*.png:
         07-home-desktop.png, 07-post-desktop.png, 07-home-mobile.png)
```

Rendered, dark console default with the brand chrome (wordmark + cursor, sidebar
`>` prompts + score block + Neko, mint `> month _` section heads with coin-slot
dividers, cartridge body at a capped measure):

![Home, desktop](../proof/07-home-desktop.png)
![Home, mobile](../proof/07-home-mobile.png)

## Negative control (revert -> RED -> restore)

The load-bearing change is `assets/scss/styles.scss` (the Arcade Neko token system).
Reverting it returns the served site to the OLD 2024 skin, captured live before this
work at `docs/brand/visuals/current/log-current.png`: white paper, thin system sans,
mint-on-every-heading, no console frame, no score block, hidden post title. That
baseline is NOT on-brand for DESIGN.md.

```
Step GREEN:   new styles.scss -> dark console, wordmark, score block, mint prompts,
              coin-slot dividers, cartridge  (07-home-desktop.png)
Step REVERT:  old styles.scss (main@cf4c338) -> white skin, no console frame, title
              hidden  (matches docs/brand/visuals/current/log-current.png, the live
              before-capture)  = RED (not the DESIGN.md brand)
Step RESTORE: new styles.scss -> dark console restored (GREEN)
Verdict: PASS  (the stylesheet is load-bearing: without it the page reverts to the
         exact pre-reskin baseline; with it the page renders the locked brand. The
         template touches in header.html / menu.html add the wordmark, breadcrumb,
         score block, and sidebar Neko that the new CSS styles.)
```

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
env -u GOROOT hugo server --bind 127.0.0.1 --port 1314
# visit http://127.0.0.1:1314/blog/log/ -> dark Arcade Neko console
# git show main:assets/scss/styles.scss > /tmp/old.scss; cp /tmp/old.scss assets/scss/styles.scss
#   -> reload: old white skin (RED). git checkout assets/scss/styles.scss -> restored (GREEN)
```

## Done-gate mapping

- Locally-served log site renders on the DESIGN.md tokens (color/type/spacing/
  components match): yes, see the token checklist in `docs/proof/07-reskin.md`.
- `docs/proof/07-reskin.md` carries before/after (desktop + mobile) + the token
  checklist: yes.
