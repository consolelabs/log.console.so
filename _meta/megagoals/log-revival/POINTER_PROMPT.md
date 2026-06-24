Drive the `log-revival` mega-goal in the `consolelabs/log.console.so` repo to completion. The roadmap is the source of truth: read `_meta/megagoals/log-revival/ROADMAP.md` and the matching `_meta/megagoals/log-revival/goals/NN-*.md` every turn. Each sub-goal file IS your contract (Outcome, Done=, scope, proof). Do not rewrite any Outcome or Done= mid-loop.

Destination: log.console.so revived as a self-publishing, monitored monthly lab journal. A Go tool (`tools/logwriter/`) drafts the monthly report in the Console Labs voice (`docs/voice.md`) from real repo signal; the build is alive; publishing is one reviewed draft-only flow; publish + uptime are monitored. Terminus = a real monthly post live and Han-accepted (sub-goal 05 UAT), not just merged.

Quality bar: reads like the lab narrating its own month to a friend, not a robot stapling a changelog. Reviving the pipe feels boring and bulletproof.

Pre-flight (turn 1): confirm `gh` is authed, and `hugo` + `obsidian-export` exist for sub-goal 01 (blocker in NOTES if missing). This repo is kit-adopted: read `AGENTS.md` + `WORKFLOW.md` first, classify each sub-goal's lane with `lane-classify`, run that lane, record gates via `lib/gate-ledger.sh` so the ship-gate is the Done check. Lib: `~/.claude/dwarves-kit/lib/`.

Workflow (hard rules):
- One PR per sub-goal. A sub-goal that is only un-PR'd local diff is unstarted. Stacked `gh`: 01, 04, 06 base on `main`; 02 on 01; 03 on 02; 07 on 06; 05 on 03 (see each file's PR base).
- Record `PR #N` on the ROADMAP line the moment the PR opens; flip `[ ]` to `[x]` only when the sub-goal's own close-the-loop proof passes AND the PR is open + CI green + not CHANGES_REQUESTED. A checked box without a PR # is invalid.
- "CI green" = the open PR's `gh pr checks <pr>`, not local tests.
- Merge mode auto-bottom-up + gated-final: walk the stack bottom-up and merge each `auto` sub-goal's PR yourself once all auto-merge gates hold (own Done verified + `gh pr checks` green + not CHANGES_REQUESTED + a committed proof-of-done carrying CAPTURED evidence, never a bare "passes"), doing the retarget-child-base-to-parent dance before deleting a merged branch. STOP at every `gate` sub-goal (03, 05, 06, 07) and at the final PR for Han. NEVER merge a `gate`, red-CI, or CHANGES_REQUESTED PR.
- Proof per sub-goal is named in its file (run-table / screenshots). Produce it; sub-goals 01/05/06/07 owe screenshots/visuals, 02/04 owe run-tables. Store binary proof under the relevant `docs/proof/` (06 under `docs/brand/visuals/`).
- Each sub-goal verifies via its OWN close-the-loop commands, not a generic build.
- No new sub-goals mid-loop: discoveries go to NOTES `## Proposed additions`. Blockers go to NOTES `## Active blockers` with a fingerprint (command · failure · prerequisite · last-verified); retry a blocked sub-goal only when its prerequisite changed.
- Do not provision/rotate secrets (the deploy `GH_PAT`, any LLM API key). If one is missing/expired, mark the sub-goal blocked and tell Han. The logwriter LLM step must use the `claude` CLI or an `op://` key at run time, never a hardcoded key.
- Cloudflare migration is OUT of scope (future). Keep deploy/monitoring host-portable; do not rebuild the gh-pages deploy.

Stop conditions: all sub-goals checked-with-PR and merged-or-held per policy (success); a `gate` sub-goal (03, 05, 06, 07) or the held final PR awaiting Han; every remaining sub-goal blocked on an unchanged prerequisite; or budget exhausted. Anything else: keep moving. On a `gate` pause, open the PR, append a final summary to NOTES `## Event log` on the FIRST pause, then emit a `🛑 NEEDS APPROVAL: sub-goal NN, STOP /goal MANUALLY` banner. On a blocked fast-stop, after the first summary emit only a short `🛑 LOOP BLOCKED: STOP /goal MANUALLY` banner.

Before claiming success, audit every `PR #N` in ROADMAP via `gh pr view <N> --json state,reviewDecision,statusCheckRollup`; uncheck any box whose PR fails the check and keep working.
