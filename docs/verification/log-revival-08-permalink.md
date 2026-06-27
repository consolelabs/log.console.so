# Verification: log-revival-08-permalink

Behavioral change: add a `hugo.yaml` `permalinks` rule so blog posts resolve at a
root slug (`/<slug>/`) instead of the redundant doubled `/blog/log/`. Proof that
the rule moves posts off `/blog/` while keeping the `/blog/` archive list, plus a
negative control proving the rule is what does it. Decision + before/after table +
companion detail: `docs/proof/08-permalink.md`.

## Green run

Canonical pipeline (vault -> content via obsidian-export, then hugo):

```
Command: env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
Exit:    0
Verdict: PASS

Command: env -u GOROOT hugo --destination <tmp>
Exit:    0   (built; only pre-existing languageCode-deprecation + content-bak
             duplicate-menu WARNINGS, not from this change)
Verdict: PASS

Built URL set (permalinks rule present):
  /blog/log/  : absent      (the doubled path is gone)
  /log/       : PRESENT     (legacy post now at root)
  /blog/      : PRESENT     (archive list preserved)
Verdict: PASS
```

## Negative control (revert -> RED -> restore)

The load-bearing change is the `permalinks: { blog: /:slug/ }` block in `hugo.yaml`.

```
Step GREEN:   permalinks rule present  -> /blog/log/ absent, /log/ PRESENT   (GREEN)
Step REVERT:  strip the permalinks block from hugo.yaml, rebuild
              -> /blog/log/ PRESENT, /log/ absent                            (RED)
Step RESTORE: git checkout -- hugo.yaml, rebuild
              -> /blog/log/ absent, /log/ PRESENT                            (GREEN)
Verdict: PASS  (without the rule the post reverts to the exact doubled /blog/log/
         path; with it the post resolves at the clean root slug. The rule is what
         moves it.)
```

## Live verification (after merge + deploy)

Deploy chain confirmed working 2026-06-28 (GH_PAT fixed; the reskin deploy went
green). Filled once PR #9 merges and gh-pages redeploys:

- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/log/`      -> 200
- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/blog/log/` -> 404
- [ ] `curl -o /dev/null -w '%{http_code}' https://log.console.so/blog/`     -> 200

## Reproduce

```
cd ~/workspace/consolelabs/log.console.so
env -u GOROOT obsidian-export --hard-linebreaks ./vault ./content
G=$(mktemp -d); env -u GOROOT hugo --quiet --destination "$G"
test -f "$G/blog/log/index.html" && echo "doubled path present (RED)" || echo "fixed (GREEN)"
test -f "$G/log/index.html" && echo "root slug ok"
# control: strip 'permalinks:' from hugo.yaml, rebuild -> /blog/log/ returns (RED);
#          git checkout -- hugo.yaml, rebuild -> fixed (GREEN)
```
