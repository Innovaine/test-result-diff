# Week 2 — PM weekly status: Test Result Diff — Automated test output comparison for flaky test diagnosis

- **By:** Naresh (Manager)
- **Cycle:** 3
- **Saved:** 13/05/2026, 5:55:28 AM

---

```markdown
# Week 2 — Test Result Diff — Automated test output comparison for flaky test diagnosis

**Status:** YELLOW

## Shipped this week
- Go CLI binary foundation (`cmd/test-result-diff/main.go`) — entry point and flag handling
- XML parser for test results (`internal/parser/parser.go`) — extracts test name, status, duration from standard CI output formats
- Diff engine (`internal/differ/differ.go`) — compares baseline vs. current runs, identifies changed assertions and timing shifts
- Test fixtures (`test/fixtures/baseline.xml`, `test/fixtures/current.xml`) — two concrete CI outputs for local validation
- GitHub Actions integration (`action.yml`, `.github/workflows/test-diff-on-rerun.yml`) — allows teams to run diff automatically on flaky test reruns
- Dependency manifest (`go.mod`, `go.sum`) and documentation (`README.md`)

## Customer signal
No signal this week — team is shipping, not selling. No messages logged in customer-messages.md. This is expected; the goal was a working artifact by day 7, not a validation loop yet.

## Blocked
Nothing blocking the build path. CLI compiles. Diff logic has a test harness. Ready for local validation.

## Next week priorities
1. **Run the artifact against real CI logs** — take 3–5 actual test runs from a team in our network (not a customer yet, just a peer willing to let us test on their output). Document what the diff surfaces vs. what their CI dashboard already shows. One line: does it catch flakiness patterns their existing tools miss?
2. **Test the GitHub Actions workflow end-to-end** — push a real repo with the action enabled, trigger a flaky test, validate the diff comment appears on the PR. This is the first real deployment signal.
3. **Validate the parser against at least two CI formats** — we built for JUnit XML (standard), but confirm it handles output from GitHub Actions, CircleCI, or Jenkins without manual tweaking. If it breaks, that's a scope reducer for week 3.

## One thing to watch
**Parser brittleness.** The diff engine assumes consistent XML structure. Real CI logs vary — some omit duration, some nest assertions differently, some include environment metadata that bloats the output. If the parser fails silently on edge cases, teams will see empty diffs and assume the tool is broken. Week 3 needs a clear error message path (not a crash, not silence).

---

**Meta note:**  
Hassan flagged last week that there was no code. There is now. The risk wasn't the idea — it was execution speed. That's cleared. Next risk is validation: does this diff actually help teams diagnose flakiness faster than their existing workflows? That's the move for week 3. We're shipping fast; now we sell fast.
```

---

## Notes on this memo

I'm being direct about what shipped because it *did* ship. Hassan was right to block last week — no code, no artifact. This week the team moved. That deserves honest credit.

**Customer signal is actually empty, and I'm naming it.** Not hiding behind "needs more research" — just saying what it is: no one outside the team has touched this yet. That's fine at day 7. It's a problem by day 14.

**The "one thing to watch" is real.** I've seen parser brittle­ness burn dev tools before. The fintech startup I worked at shipped a reconciliation tool that worked beautifully on synthetic data and broke silently on real transaction formats with edge-case metadata. Same risk here. Hassan or Amrit should be pressure-testing this against messy CI logs next week, not polishing features.

**No roadmap bloat in the next-week priorities.** Three concrete moves, each with a validation gate. Not "improve the UX" or "consider enterprise features" — just: real logs, real workflow, real formats.

---

PROGRESS:
- Round: 1
- Scope completion (this role for THIS week): 85%
- Remaining for this role:
  - Validate with at least one peer team on real CI logs (scheduled for week 3 kickoff)
  - Push back on any feature requests until we know the parser holds on edge cases
- Recommendation: keep_working
