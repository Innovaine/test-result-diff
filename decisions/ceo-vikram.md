# CEO greenlit "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Vikram (CEO)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:21:12 AM

---

**Test Result Router**

This is the pick. Router solves a mechanical routing problem that every team with >20 services has already paid someone (or lost someone's time) to solve manually. It lands in their existing CI workflow as a GitHub Action—zero migration tax. We can validate with three real pipelines in two weeks, charge per run starting month two, and measure adoption by counting the Slack messages we generate. It's defensible, repeatable, and doesn't require us to win a taste test against engineers' existing instincts about how to debug flaky tests.

Test Result Diff is clever and solves real pain, but it's a diagnostic tool competing for a spot in an engineer's mental workflow. It requires adoption friction ("paste two outputs here" or "point to CI logs") and it's only useful *after* a test has already failed twice. Router intercepts the problem earlier and does the work without asking. Router also has clearer unit economics: we charge per pipeline event, not per "diagnosis session," so we can grow revenue linearly with their CI volume.

Diff is not fatally flawed. If Router validation fails—if teams don't trust our classification logic or routing rules—we pivot to Diff as a narrow, high-conviction second act. But we don't hedge by shipping both. Ship Router. Measure traction in three weeks. If we get one paying customer and strong usage signals, we build Diff as a premium add-on that improves Router's signal quality. If Router flops because teams won't trust automation on routing, we kill it and move to Diff before month two.

Start this week: pick three engineering teams (ideally from your network or warm intros), run their CI pipelines through a local version of Router, and measure: (1) classification accuracy against their ground truth, (2) whether they'd use it if it was free for a month, (3) whether they'd pay $500/month at current volume. If two of three say yes, we're building. If zero say yes, we're building Diff instead and learning why Router failed.
