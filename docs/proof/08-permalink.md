# Proof: sub-goal 08, permalink slug fix

**Done =** `hugo.yaml` declares the post permalink, a built post resolves at the agreed clean slug, the redundant `/blog/log/` post duplicate is gone, and this doc carries the before/after table + a live check.

## Decision (Han, 2026-06-28)

Canonical post URL = **drop the redundant `/blog/` prefix, posts at a root slug** (`/<slug>/`). `log.console.so` already says "log", so `/blog/` doubled the concept (the legacy `blog/log.md` rendered at `/blog/log/`). The section archive list stays at `/blog/`.

Implemented as a one-line Hugo rule:

```yaml
permalinks:
  blog: /:slug/
```

## Before / after

| Post (vault file) | Before | After |
|---|---|---|
| `blog/log.md` (legacy 2023 changelog) | `/blog/log/` | `/log/` |
| `blog/shape-up.md` | `/blog/shape-up/` | `/shape-up/` |
| `blog/console-log-2026-06.md` (first monthly, in content PR #4) | `/blog/console-log-2026-06/` | `/console-log-2026-06/` |
| section archive list | `/blog/` | `/blog/` (unchanged) |

## Confirmation run-table (local build)

`hugo --destination <tmp>` over the live vault submodule:

| Check | Expected | Result |
|---|---|---|
| `/blog/log/` (the doubled path) gone | absent | `NO (fixed)` |
| `/log/` (legacy post at root) | present | `YES` |
| `/shape-up/` (post at root) | present | `YES` |
| `/blog/` (archive list) preserved | present | `YES` |

Negative control: before adding the `permalinks` block, the same build produced `/blog/log/` and `/blog/shape-up/` and NO `/log/` , confirming the rule is what moved them.

## Live verification (after merge + deploy)

Filled once this PR merges and the gh-pages deploy runs (deploy chain confirmed working 2026-06-28 after the GH_PAT fix):

- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/log/` -> 200
- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/blog/log/` -> 404 (doubled path gone)
- [ ] `https://log.console.so/blog/` -> 200 (archive list intact)

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
D=$(mktemp -d); hugo --quiet --destination "$D"
test -f "$D/blog/log/index.html" && echo "doubled path still here (bad)" || echo "fixed"
test -f "$D/log/index.html" && echo "root slug ok"
```

## Out of scope (logged to NOTES Proposed additions)

The build also publishes a `content-bak/` backup tree from the vault (`/content-bak/blog/log/`, `/content-bak/earn/*`, ...) , a stale duplicate of the whole site. That is a vault-content cleanup (unpublish/remove `content-bak/` in `consolelabs/content`), not a permalink rule, so it is filed as a discovery, not fixed here.
