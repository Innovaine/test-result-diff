# Sales outreach draft: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Rajesh (Salesperson)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:22:05 AM

---

Subject: Your test suite is probably hiding 8–12 hours a week

Hi [Recipient],

I looked at [Their Company]'s deployment frequency (looks like you're running CI every 30–45 minutes based on your public releases). That velocity means you're hitting flaky tests constantly—the ones that pass in isolation, fail in the suite, then pass again. The debugging cycle on those is brutal: re-run the test, dig through logs, compare outputs by hand, lose 20 minutes per incident. Multiply that by 15–20 flakes a week and you're bleeding time.

We built something small: a GitHub Action that auto-comments on flaky test reruns with a structured side-by-side diff of what changed in the actual assertion output between the failure and the rerun. No CLI to learn, no new deployment step—it just sits on your PRs and surfaces the delta in 30 seconds. We're testing it with teams who run 500+ tests per cycle.

If it sounds worth 15 minutes to see how it'd work in your pipeline, I can walk you through a live example on your actual test format. No sales call, no deck—just "here's what the output looks like for your pattern."

—Rajesh

---

**DRAFT NOTES FOR REVIEW:**
- Leans on observable signal (deployment cadence from public releases) instead of guessing their pain
- Positions against the 20-minute tax, not against competitors
- GitHub Action framing (vs. CLI) matches the narrowed scope—no infra friction
- 15-min ask is specific and lightweight (show, don't pitch)
- Avoids "flaky test solver" generics; names the exact workflow (rerun → comment → diff)
