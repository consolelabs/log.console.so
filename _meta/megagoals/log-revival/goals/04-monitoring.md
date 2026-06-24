# Sub-goal 04: monitoring

**Merge policy:** auto
**Time budget:** 2-4 hours of loop work
**Proof:** run-table / catalog output from vps-mon showing the new checks as `monitored` (not `gap`), plus the check definitions committed
**Depends on:** 01
**Branch:** feat/log-revival-04-monitor
**PR base:** chore/log-revival-01-build

## Outcome

log.console.so's health is monitored through Han's existing vps-mon, in three signals: the site is **up** (HTTP 200 on the live URL), the content is **fresh** (a post published in the current or prior month, so a missed monthly log surfaces), and the **publish job** ran (the deploy Action's last run is green). Checks are host-portable: uptime + freshness survive the future Cloudflare move untouched; only the deploy-job signal is host-specific.

## Quality bar

If the site goes down, or a month slips by with no log published, Han finds out from monitoring, not by remembering to look. Wiring follows the established vps-mon onboarding so it shows up in the catalog like every other monitored job, no bespoke alerting.

## How to close the loop

- Follow the `job-monitoring-onboarding` runbook: classify each signal's reconcile kind (uptime = on-demand/probe, freshness = scheduled, publish-job = scheduled/daemon) and route it to the right vps-mon liveness signal.
- Add the check definitions where vps-mon expects them; run the vps-mon catalog and capture output showing the log.console.so checks as `monitored` into `docs/proof/04-monitoring.md`.
- The freshness check reads the latest published post's date (from the live site's RSS/JSON feed or the content repo) and flags if older than ~5 weeks.

**Done =** vps-mon's catalog lists the log.console.so uptime, freshness, and publish-job checks as `monitored` (not `gap`), the check definitions are committed, and `docs/proof/04-monitoring.md` carries the catalog output.

## Scope edges

**In:** vps-mon check definitions for uptime + freshness + publish-job, the onboarding wiring, the proof capture.
**Out:** the writing tool (02), the publish flow (03), building a new monitoring system.
**Not:** standing up a new monitoring stack, alerting channels beyond vps-mon's existing ones, monitoring the content authoring, Cloudflare-specific deploy hooks (revisit at the CF move).

## Where to look

The `vps-mon` tool in ops-toolkit (`~/workspace/tieubao/ops-toolkit/tools/vps-mon/`) and its catalog/reconcile model, the `job-monitoring-onboarding` skill runbook, the live site's feed (`/index.json` or RSS) for freshness, the deploy Action for the publish-job signal.

## PR body

Wire log.console.so into vps-mon: uptime (HTTP 200), monthly freshness (a post in the current/prior month), and publish-job (deploy Action green). Host-portable ahead of the Cloudflare move. Follows the job-monitoring-onboarding runbook.

Verify: vps-mon catalog shows the three checks as `monitored`; see `docs/proof/04-monitoring.md`.

Stacked on #<PR-01>; review after it. Part of mega-goal `log-revival`.

## Notes
