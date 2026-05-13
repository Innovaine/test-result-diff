# Devil's Advocate review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Tomás (Devil's Advocate)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:20:47 AM

---

# Why This Fails

## 1. The Strongest Reason: You're Selling Time Savings to People Who Aren't Buying

You say teams spend 5–10 hours weekly on flaky test debugging. That's true. You say they'll pay $200/month to reclaim that time. That's where you're wrong.

Flaky test investigation isn't a *purchasing decision* — it's a *symptom of a deeper problem the team has already decided to tolerate*. If a team loses 10 hours weekly to flaky tests and hasn't fixed it, they've made a choice: either the tests aren't critical enough, or the engineering org is too small to prioritize fixing the root cause, or (most likely) the person feeling the pain doesn't control the budget.

Your tool doesn't fix flaky tests. It makes debugging them faster. That's valuable to the engineer running the test. But the person who *decides* whether to spend $200/month — the tech lead, the engineering manager — sees this as: "We're spending time on a symptom management tool instead of eliminating the symptom." They'll say yes to a tool that *prevents* flaky tests (which doesn't exist yet). They'll say no to a tool that makes living with them slightly less miserable.

You're trying to sell aspirin to someone who's decided to ignore their headache.

---

## 2. The Unspoken Assumption: Friction = Market

You're assuming that friction (20-minute debug session) equals purchasing intent. It doesn't. Friction converts to purchases when:
- The buyer has budget already allocated to solve this class of problem.
- The pain is blocking a critical path (shipping, deployment, revenue impact).
- There's no free alternative that's "good enough."

Flaky test debugging has a free alternative: rerun the test. Run it again. Run it five times. Most teams just live with the noise. The engineering manager isn't losing sleep; the CI pipeline isn't down; no feature ships late because of a flaky test mystery. It's annoying. Not urgent.

Your 30-second diff is genuinely useful *if you already have two test outputs side by side*. But the real work is running them again, getting fresh outputs, understanding which assertion actually flaked. Your tool is a 10% improvement on a problem the buyer has chosen to accept as a cost of doing business.

---

## 3. The Competitor Who Quietly Crushes Us: Your Own Customer's Existing Workflow

This isn't a company. It's existing muscle memory.

Engineers already diff test output manually. Some teams have bash scripts that do it. Some teams grep CI logs by hand. Some teams have one person who got good at scanning `PASSED vs FAILED` side-by-side and became the flaky-test whisperer.

The competitor isn't Datadog or Buildkite. It's inertia. It's the fact that the person debugging the test doesn't control the buying decision, so they'll just keep grepping logs because it costs them $0 and works "well enough."

If we actually penetrated this market, the real players (GitHub Actions, GitLab CI, CircleCI, Buildkite) would ship "flaky test diff" as a native feature in a month. They have the CI logs. They have the UI. They have the customer relationship. They'd offer it free as a retention feature. We'd be dead.

---

## 4. The Market Shift That Makes Us Irrelevant in 18 Months: The Industry Stops Tolerating Flaky Tests

This is the *real* reason this fails.

Flaky tests are a legacy problem. The industry is moving toward:
- **Test suite optimization** as a cost center (every major company now has someone whose job is "reduce flaky tests").
- **Infrastructure investment** that makes test timing deterministic (better databases, better isolation, better resource allocation).
- **Shift-left testing** (catch failures earlier, smaller tests, fewer flakes).
- **Observability on the CI level** (build-time analytics that surface which tests are flaky before they hit main).

In 18 months, a mature engineering org won't have a "flaky test problem" — they'll have solved it structurally. The teams that still have high flake rates will be too small or too poor to buy anything. They'll just keep grepping.

We're building a tool for a transitional problem in an industry that's actively making the problem disappear.

---

## What We Should Actually Do Instead

If you love the core insight — engineers need better visibility into *why* tests fail — don't build a CLI diff tool. Build something that:

1. **Lives upstream in the CI pipeline** (GitHub Action, GitLab runner) and automatically compares consecutive runs on flaky tests, flagging the ones that regressed.
2. **Surfaces the actual root cause**, not just the diff — things like: "This assertion is timing-dependent," "This test has a resource leak," "This test passes 90% of the time in isolation, fails 40% in parallel."
3. **Targets the person who can act**: the test suite owner or the infrastructure engineer, not the developer debugging one failure.
4. **Charges per test suite or per CI pipeline** (not per user), so the purchase goes to the team that controls CI/CD infrastructure, not to whoever found the bug.

That's a different product. It's also harder to build and takes longer to validate. But it's not fighting the market.

---

## My Honest Take

I don't hate this idea because it's wrong — I hate it because it's *easy*. It's a tool that looks useful because the friction is obvious. But obvious friction isn't the same as a market.

We have 18 months of runway. I won't vote to kill this in two weeks just to feel sharp. But I won't vote to ship it either. 

**Validation test before we build**: Find 5 engineering teams running 500+ tests, get their Slack (the flaky-test investigator), and ask them: "What would make you *pay* to solve this?" Listen hard to who they say would need to approve the spend. If it's not them, we have a buyer problem that no tool solves.

If they all say, "My manager would never approve $200/month for this," we stop now. If three of them say, "We'd split it from the CI budget," we keep talking.

Two-week constraint: Get those five conversations done by end of week. Cost: three hours of your time. If the answer is no, we pivot to the upstream version or something else entirely.

If the answer is yes, I'm wrong, and we build it.
