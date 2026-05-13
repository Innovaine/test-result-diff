# Chairman approved: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Amara (Chairman)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:21:22 AM

---

APPROVED WITH PIVOT: Ship as a GitHub Action that auto-comments on flaky test reruns with a structured diff of assertion output deltas — no CLI, no multi-repo sync yet.

This is the right problem (flaky tests are real pain, no workaround, monetizable), but the CLI-first wedge has friction we can't afford pre-revenue. Engineers won't install a new CLI tool for one task when they're already context-switched; they *will* see a comment on a PR that just shows them what changed between runs. GitHub Action gets us into their workflow *right now*, produces real usage signal fast, and lets us learn which test types actually create debug friction before we build the CLI.

Ship in 10 days: parse reruns from CI logs, extract assertion diffs, post structured comment. Real teams, real flaky tests, real signal on whether they'd pay. If it works, the CLI becomes the next cycle's pivot; if it doesn't, we've spent two weeks not a month on a hunch.
