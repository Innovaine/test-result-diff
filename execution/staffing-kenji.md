# HR staffing plan: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Kenji (HR)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:21:50 AM

---

# TEAM STAFFING: Test Result Diff (GitHub Action)

**Organizational home:** Extend Engineering (existing team).

**Rationale:** This is a single-product sprint, tight scope, shipping artifact is code (our primary output). No new department needed. Engineering owns the build, we don't have a separate DevTools or Integrations department yet.

---

## TEAM COMPOSITION (5 people, 3-week cycle)

1. **1 Lead Engineer (full-cycle)** — Owns GitHub Actions API integration, assertion diff algorithm, output formatting. Must have shipped a GitHub App or Action before (not negotiable—learning the platform during execution will kill the timeline).

2. **1 Backend Engineer** — Builds the diff-comparison service. Takes two test output blobs, returns structured delta (which assertions changed, what values shifted). This is the core logic; can be a small service or inline, but needs someone who can think in data transformations.

3. **1 Frontend/UX Engineer** — Designs the comment template on the PR. GitHub has constraints; the output needs to be scannable in 10 seconds and clickable (links to full diffs, logs). This is not design theater—it's the user-facing artifact.

4. **1 QA/Integration Engineer** — Tests against real GitHub repos with real flaky test suites (we need 2-3 test repos to validate against during the sprint). Catches the "works in isolation, breaks on real CI" problem.

5. **1 Product/Definition (0.5 FTE, shared with founder or existing PM)** — Owns the acceptance criteria and the "ship/no-ship" call at day 14. Not running code, running the gate.

---

## EXECUTION MILESTONES (7-14 days to ship)

- **Days 1-2:** GitHub Action scaffold, test repo setup, assertion diff algorithm drafted.
- **Days 3-5:** Service integration, comment template design (iterate on 2-3 real repos).
- **Days 6-10:** Edge cases (multi-language test output, large diffs, timeout handling).
- **Days 11-13:** Real repo testing, bug fixes, documentation.
- **Day 14:** Ship to marketplace + 3-5 beta installs (warm handoff, measure time-to-first-value).

---

## CONCERNS (raised after delivery, as required)

**Scope drift risk:** The Chairman's pivot from CLI to GitHub Action is smart (lower friction, immediate value signal), but "auto-comment on flaky test reruns" assumes we can reliably detect which reruns are flaky vs. which are legitimate failures. If the heuristic is wrong, we'll spam PRs and get uninstalled. Need to spec this hard on day 1: **What triggers the action?** (only on retry, only on branch, only if test name appears twice with different outcomes?). If that's fuzzy, we ship with a **manual trigger only** (engineer pastes two outputs in the issue, action runs), which is less magical but prevents false positives.

**Still concerned about**: Flaky test market size at pre-revenue. We're betting $50k+ in engineering time that teams will pay $200/month for this. That's an untested assumption. After day 14, before we scale, we need **3 hard conversations with engineering teams running 500+ tests** (Slack, Twitter, direct outreach—no deck, just "does this solve your problem?"). If fewer than 2 say "yes, we'd pay," we shelve it and rotate to the next idea. Don't let momentum carry us into a dead end.

**Recommend:** Add a "Day 15 validator" task—someone reaches out to 5 beta users, measures: (1) install-to-first-use time, (2) whether they re-run it without being prompted, (3) would they pay. That's your real pass/fail gate, not feature completeness.
