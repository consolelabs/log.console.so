# NOTES, log-revival

## Active blockers

[blocked] sub-goal 04 (monitoring) · symptom: wiring log.console.so into vps-mon · failure: (1) the uptime HTTP check requires inserting a row into the LIVE vps-mon D1 on Han's Cloudflare account (outward infra mutation); (2) the requested "freshness" + "publish-job-success" signals are NOT natively supported by vps-mon (no content-freshness check; a GitHub Action is not a vps-mon discovery source) and need new Worker code; (3) the monitoring config belongs in ops-toolkit/tools/vps-mon, not this repo (cross-repo). · prerequisite: Han runs/approves the vps-mon D1 change + decides whether to extend vps-mon for freshness/publish-job, OR rescopes 04 to uptime-only · last verified: 2026-06-25

## Proposed additions

- 2026-06-25: future mega-goal idea, migrate log.console.so from GitHub Pages to Cloudflare Pages, Han wants it later. Revisit the host-specific parts of sub-goals 03 (deploy chain) and 04 (publish-job signal) at that point. Sub-goals here were kept host-portable so this move stays cheap.
- 2026-06-25: vps-mon extension (for sub-goal 04), add a content-freshness check type (read a site's feed date, alert if no post in N weeks) + a GitHub-Action-success signal (deploy pings a heartbeat on success). Both need new vps-mon Worker code in ops-toolkit/tools/vps-mon. Until then, 04 can only do the uptime HTTP check (which itself needs a live D1 insert on Han's CF account).
- 2026-06-28 (Han): permalink slug fix. The post URL renders as `log.console.so/blog/log/`, which is wrong and redundant: "blog" + "log" doubles the same concept on a site that IS the log, and `/blog/log/` is the section page (title "Log | Console Labs"), not a per-post slug derived from the filename (`blog/console-log-2026-06.md` should be `/blog/console-log-2026-06/` or similar). Cause: no `permalinks:` block in hugo.yaml and no `slug` in the post front-matter, so the section/URL structure is doing the wrong thing. Fix scope: decide the canonical post URL (drop the redundant `blog/` prefix, or use a dated slug), set it via a hugo.yaml `permalinks:` rule + optional front-matter `slug`, and verify no stale `/blog/log/` duplicate remains. Belongs with the publish/reskin surface (sub-goals 03/07); verify on the live deploy once the GH_PAT blocker is cleared. → PROMOTED to a formal sub-goal `08-permalink-fix` on Han's call (2026-06-28).
- 2026-06-28 (discovered during 08): the vault publishes a `content-bak/` backup tree (`/content-bak/blog/log/`, `/content-bak/earn/*`, a full stale copy of the site). It is live and a real duplicate surface. Fix = unpublish/remove `content-bak/` in the `consolelabs/content` vault repo (or add `_build: {render: never}` / move it out of the content root). Vault-content cleanup, not a permalink rule, so out of 08's scope. Candidate for a small follow-up sub-goal or a direct vault PR.
- 2026-06-28 (Han, future sub-goal `de-obsidian`): Han no longer uses the Obsidian app, so the Obsidian build coupling is now pure overhead worth removing eventually. NOT forced , `obsidian-export` is a build-time Rust CLI in CI + devbox, not the app, so the site builds headless without anyone opening Obsidian, and new `logwriter` posts are plain markdown (no new coupling). But a full de-Obsidian is real work: the existing content uses 167 wikilinks `[[...]]` (31 files) + 27 embeds `![[...]]` + 7 templater `<% %>` blocks that `obsidian-export` resolves at build time; Hugo can't render them natively. Scope when done: convert wikilinks -> Hugo `ref`/`relref`, embeds -> shortcodes, resolve templater; drop `obsidian-export` from Makefile + `.github/workflows/main.yml` + `devbox.json`; remove `vault/.obsidian/` + `vault/_templater/`; point Hugo at the content directly. Medium effort, regression risk on the 31 legacy posts. No forcing function; do it when simplifying.
- 2026-06-28 (runbook gap, fixed in the slug-basename PR): changing a published post's slug breaks its old URL. The publish runbook now documents adding `aliases: ["/old-url/"]` to the post front-matter so Hugo emits a redirect from the old path.

## Event log

2026-06-25 HH:MM · scaffolded · mega-goal log-revival created (5 sub-goals); kit adopted + repo context (CLAUDE.md, docs/voice.md) committed on branch chore/kit-adopt-context.
2026-06-25 HH:MM · scope added · design track added per Han: 06-design-system (DESIGN.md brand guideline via UI/UX designer sub-agent) + 07-reskin-log (apply brand to log layouts).
2026-06-25 HH:MM · autonomy change · Han: remove all gates, review at the end, presume the loop's call. All sub-goals now `auto`; only 05 (first post) is the held-final. 06 reworked into a rigorous multi-step design process (audit, parallel divergent directions, /visual-team critique, converge, mockup) since there is no mid-loop human gate. 05 now fan-in on 02/03/07 so the first public post debuts on the new brand.
2026-06-25 HH:MM · foundation merged · PR #1 (kit adopt + context + scaffold) to main.
2026-06-25 HH:MM · sub-goal complete · 01-revive-build, PR #2, merged cf4c338. Build revived on Hugo 0.163.3 (was 0.120.2); stack upgraded to latest per Han; obsidian-export rebuilt; CI de-staled; build-check.yml added as the mega-goal's PR CI. Negative-control proof.
2026-06-25 HH:MM · sub-goal complete · 02-writing-tool, PR #3, merged 33f2c5f. Go tools/logwriter gathers a month's merged-PR signal (gh) + drafts in-voice (Anthropic API default, claude CLI fallback). Proven on June 2026 (155 PRs). Privacy-trimmed sample.

==== CHECKPOINT, autonomous run, 2026-06-25 ====
Done: foundation (PR #1), 01-revive-build (PR #2), 02-writing-tool (PR #3). All merged with proof + negative controls.
Blocked: 04-monitoring (see ## Active blockers, needs Han's vps-mon CF infra + extensions).
Remaining + their real prerequisites:
- 06-design-system: unblocked, large (multi-agent design). On the critical path to 05.
- 07-reskin-log: depends on 06.
- 03-publish-flow: needs write access to consolelabs/content + the live deploy path (GH_PAT for the private-submodule checkout is unconfirmed; `gh secret list` returned empty).
- 05-first-post-uat: the held-final (Han's end-review); depends on 02/03/07.
Decision points for Han: (a) rescope or run 04's vps-mon wiring; (b) confirm the GH_PAT deploy secret for 03/05; (c) whether to run the design track (06/07) now or in a focused session.
================================================

2026-06-28 · GH_PAT deploy secret FIXED (Han minted a fine-grained `github-console-deployer-token`, set as the consolelabs org `GH_PAT`). The "Deploy Hugo from Obsidian notes" workflow now passes; the Arcade Neko reskin (07) is finally LIVE. This clears the deploy blocker on 03/05.
2026-06-28 · sub-goal complete · 08-permalink-fix, PR #9, merged 4047665. `hugo.yaml` permalinks `blog: /:slug/` drops the redundant `/blog/` prefix; the doubled `/blog/log/` is gone (404 live), posts at root (`/log/`, `/shape-up/` 200), `/blog/` archive intact. Green + negative control in docs/verification/log-revival-08-permalink.md.
2026-06-28 · STATE · build stack 01,02,03,06,07,08 merged + live. Remaining: 04-monitoring (blocked on vps-mon CF infra, unchanged) and 05-first-post-uat (held-final, PR #8 open; its deploy blocker is now cleared, so Han's end-review is unblocked: merge consolelabs/content PR #4 to publish the first post).
