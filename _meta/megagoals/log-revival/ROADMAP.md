# Mega-goal: log-revival

**Destination:** log.console.so is revived as a self-publishing, monitored monthly lab journal: a Go writing tool drafts the monthly report in the Console Labs voice from real repo signal, the build pipeline is alive, publishing is one reviewed draft-only flow, and publish + uptime are monitored. Terminus = a real monthly post live and Han-accepted (UAT), not just merged.
**Quality bar:** Reads like the lab narrating its own month to a friend, not a robot stapling a changelog. Reviving the pipe feels boring and bulletproof: nobody fights Hugo, a stale submodule, or a missed-month again.
**Stacking tool:** gh (stacked PRs)
**Merge mode:** auto-bottom-up
**Merge autonomy:** gated-final
**Started:** 2026-06-25

## Sub-goals

- [x] 01-revive-build, local `make build` produces `public/` on the current toolchain (versions pinned), `auto`, PR #2
- [x] 02-writing-tool, `logwriter <month>` emits a vault-ready in-voice draft from real repo signal, `auto`, PR #3
- [ ] 03-publish-flow, one-command draft-only flow stages a post into the content repo + opens a PR (no direct push to live), `auto`, PR #
- [ ] 04-monitoring, vps-mon catalog shows log.console.so uptime + freshness + publish-job as `monitored`, `auto`, PR #
- [ ] 06-design-system, `docs/brand/DESIGN.md` brand guideline locked via a multi-step process (audit, parallel divergent directions, critique, converge, mockup), `auto`, PR #
- [ ] 07-reskin-log, log Hugo layouts + assets render on the DESIGN.md tokens, `auto`, PR #
- [ ] 05-first-post-uat, a real monthly post is live + Han-signed-off, `auto` (HELD-FINAL: opened then held for Han's end-review), PR #

## Dependencies

- 02 depends on 01 (needs a working build to verify the draft renders)
- 03 depends on 02 (needs a draft to publish)
- 04 depends on 01 (the deploy job + live site to monitor must exist)
- 06 depends on nothing (reads the current sticker + log styling; can start immediately)
- 07 depends on 06 and 01 (implements the locked brand on a working build)
- 05 depends on 02, 03, and 07 (the first public post debuts on the new brand); fan-in, opened after all three merge
- **Held-final:** 05 is the one PR the loop does NOT merge. Everything else is `auto` and the loop merges it bottom-up (including the live-deploying reskin). The loop builds + rebrands the site autonomously, then opens 05 and stops; Han's single end-review = merging 05, which publishes the first real post. This is "remove all gates, I review at the end."

## Assumptions (baked from Han's 2026-06-25 answers)

- **Cadence: monthly.** One self-narration report per month; the tool is not built for weekly output.
- **Tool depth: hybrid.** logwriter assembles signal AND drafts each section in-voice, section-regeneratable; the human edits.
- **Publish: draft-only mechanism.** The publish flow opens a content PR rather than pushing straight to live; the first real post (05) is the held-final, so no actual post publishes without Han's merge.
- **Autonomy (Han, 2026-06-25 second directive): no gates, all `auto`, gated-final.** The loop makes every call itself and merges every sub-goal bottom-up EXCEPT the held-final (05). Han reviews once, at the end, by merging 05. Consequence Han accepted: the tooling and the reskin auto-merge, and the reskin auto-deploys live before that end-review; only the first-post publish is held.
- **Tool home/lang: Go, `tools/logwriter/` in this repo.** Host-agnostic, survives the future Cloudflare move.
- **Cloudflare move is a future follow-on, NOT in this mega-goal.** Sub-goals stay host-portable (uptime/freshness checks, the content-vault publish flow); the existing gh-pages deploy is only lightly un-rotted, not rebuilt. See NOTES `## Proposed additions`.
- **Kit-adopted repo:** sub-goals route through the dwarves-kit lane (read AGENTS.md + WORKFLOW.md, classify with lane-classify, record gates via gate-ledger.sh).

## Audit cheat sheet

    grep -oE 'PR #[0-9]+' ROADMAP.md | sort -u | while read -r _ pr; do
      gh pr view "$pr" --json state,reviewDecision,statusCheckRollup
    done
