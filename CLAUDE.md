# log.console.so

Console Labs' public lab journal. This is where the lab narrates itself: monthly
reports, release notes, thought-process, and experiment write-ups, published in
the Console Labs voice. Live at [log.console.so](https://log.console.so).

## Architecture (read before changing anything)

Two repos, one site:

| Piece | Repo | Role |
|---|---|---|
| Site engine | `consolelabs/log.console.so` (this repo) | Hugo + custom `layouts/`, build pipeline, deploy |
| Content | `consolelabs/content` (the `vault/` submodule) | an Obsidian vault; the actual posts live here |

**Publish pipeline:** Obsidian vault (`vault/`) gets converted to Hugo content by
`obsidian-export`, Hugo builds `public/`, and `public/` deploys to the `gh-pages`
branch (CNAME `log.console.so`).

```
write note in consolelabs/content (Obsidian vault)
  -> obsidian-export --hard-linebreaks ./vault ./content   (Obsidian md -> Hugo md)
  -> hugo -DEF --minify                                     (build to ./public)
  -> peaceiris/actions-gh-pages -> gh-pages branch          (GitHub Pages)
```

- **Deploy trigger:** `.github/workflows/main.yml` runs on push to `main`. Built-in
  `GITHUB_TOKEN` does the Pages deploy; a `GH_PAT` secret checks out the private
  content submodule.
- **Pull new content:** `.github/workflows/dispatch.yml` (manual `workflow_dispatch`)
  bumps the `vault/` submodule to the latest `content` commit, then the push to
  `main` re-triggers `main.yml`. So content edits land in `consolelabs/content`
  first, then a dispatch pulls them into the site.
- **Local preview:** `make watch-run` (or `devbox run watch-run`). Needs `hugo`,
  `obsidian-export` (cargo), and the `vault/` submodule cloned.

## Content sections (in the vault)

`blog/` (posts + the monthly log), `experiments/`, `playbook/` (engineering
practices), `earn/` (bounties), `members/` (author profiles). Identity docs at the
root: `purpose.md`, `philosophy.md`, `about.md`. New posts use the Obsidian
Templater stubs in `vault/_templater/` (`New Post.md`, `New Earn.md`, `New Member.md`).
The `Refresh *` templater scripts regenerate index/earn/member tables via Dataview.

## The voice (self-narration)

Console Labs writes in first-person-plural as a small, finance-obsessed, gamer-at-heart
lab that is a sibling to Dwarves Foundation. Casual, fun, "vibes" energy, emoji-friendly,
web3-native, honest about what failed. Author persona is usually `neko-san`. The four
pillars: **Small but Sharp**, **Financial-Centric**, **Accessible**, **Gamified**.
Full voice spec + a monthly-report skeleton: `docs/voice.md`. Match it; do not write
in a corporate or press-release register.

## Dormancy note

The site was last published Feb 2024. It is being revived as the monthly lab-report
channel. Treat existing posts as voice reference, not as current state.

<!-- kit:adopt -->
## Operating layer (dwarves-kit)

@AGENTS.md

Before touching code, classify the lane: `bash /Users/tieubao/.claude/dwarves-kit/lib/lane-classify.sh classify "<task>"`.
A full-lane change records its gates via `/Users/tieubao/.claude/dwarves-kit/lib/gate-ledger.sh` or the ship-gate blocks the push.
<!-- /kit:adopt -->
