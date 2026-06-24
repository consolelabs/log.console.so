# Mega-goal: log-revival

**Destination:** log.console.so is revived as a self-publishing, monitored monthly lab journal: a Go writing tool drafts the monthly report in the Console Labs voice from real repo signal, the build pipeline is alive, publishing is one reviewed draft-only flow, and publish + uptime are monitored. Terminus = a real monthly post live and Han-accepted (UAT), not just merged.
**Quality bar:** Reads like the lab narrating its own month to a friend, not a robot stapling a changelog. Reviving the pipe feels boring and bulletproof: nobody fights Hugo, a stale submodule, or a missed-month again.
**Stacking tool:** gh (stacked PRs)
**Merge mode:** auto-bottom-up
**Merge autonomy:** gated-final
**Started:** 2026-06-25

## Sub-goals

- [ ] 01-revive-build, local `make build` produces `public/` on the current toolchain (versions pinned), `auto`, PR #
- [ ] 02-writing-tool, `logwriter <month>` emits a vault-ready in-voice draft from real repo signal, `auto`, PR #
- [ ] 03-publish-flow, one-command draft-only flow stages a post into the content repo + opens a PR (no direct push to live), `gate`, PR #
- [ ] 04-monitoring, vps-mon catalog shows log.console.so uptime + freshness + publish-job as `monitored`, `auto`, PR #
- [ ] 05-first-post-uat, a real monthly post is live + Han-signed-off, `gate`, PR #
- [ ] 06-design-system, `docs/brand/DESIGN.md` brand guideline locked (via UI/UX designer sub-agent), `gate`, PR #
- [ ] 07-reskin-log, log Hugo layouts + assets render on the DESIGN.md tokens, `gate`, PR #

## Dependencies

- 02 depends on 01 (needs a working build to verify the draft renders)
- 03 depends on 02 (needs a draft to publish)
- 04 depends on 01 (the deploy job + live site to monitor must exist)
- 05 depends on 02 and 03 (runs the tool + the publish flow end to end)
- 06 depends on nothing (reads the current sticker + log styling; can start immediately)
- 07 depends on 06 and 01 (implements the locked brand on a working build)
- **Merge order for a rebranded launch:** 05 and 07 are both Han-merged gates. 05 does NOT hard-depend on 07 (a first post can ship on the current skin), but merge 07 before 05 if the first public post should land on the new brand (recommended). Han controls both gate merges, so this is a sequencing choice at merge time, not a code dependency.

## Assumptions (baked from Han's 2026-06-25 answers)

- **Cadence: monthly.** One self-narration report per month; the tool is not built for weekly output.
- **Tool depth: hybrid.** logwriter assembles signal AND drafts each section in-voice, section-regeneratable; the human edits.
- **Publish: draft-only.** The tool/flow prepares + opens a content PR; nothing goes live without Han's merge.
- **Tool home/lang: Go, `tools/logwriter/` in this repo.** Host-agnostic, survives the future Cloudflare move.
- **Cloudflare move is a future follow-on, NOT in this mega-goal.** Sub-goals stay host-portable (uptime/freshness checks, the content-vault publish flow); the existing gh-pages deploy is only lightly un-rotted, not rebuilt. See NOTES `## Proposed additions`.
- **Kit-adopted repo:** sub-goals route through the dwarves-kit lane (read AGENTS.md + WORKFLOW.md, classify with lane-classify, record gates via gate-ledger.sh).

## Audit cheat sheet

    grep -oE 'PR #[0-9]+' ROADMAP.md | sort -u | while read -r _ pr; do
      gh pr view "$pr" --json state,reviewDecision,statusCheckRollup
    done
