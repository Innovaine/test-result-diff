# CFO review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Lucia (CFO)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:18:32 AM

---

# Test Result Diff — Financial Reality Check

I'm going to walk through this with real numbers. This idea has merit—flaky tests are a genuine cost center—but I'm seeing gaps in the model that matter before we spend runway on it.

---

## 1. PRICING MODEL — COMMITTED NUMBER

**$299/month, billed annually.**

Here's my reasoning:
- Target customer (500+ tests/deploy) runs maybe 20–30 deploys per week. A 20-minute flaky test investigation happens 2–3 times weekly = 40–60 minutes of senior engineer time lost per week. At $150/hr loaded cost, that's $100–150 weekly friction per team.
- $299/month = $70/week saved. That's in-the-noise territory for most teams. Too cheap to signal value.
- Competitors in adjacent space (test analytics, observability add-ons) charge $400–600/month. We're underpriced if anything.
- Annual billing locks in cash flow and reduces churn risk—critical for a bootstrapped company.

**Tier structure:** Single tier. No "Pro" or "Enterprise" nonsense until we have 10 paying customers and proof that teams actually want additional features. One SKU, one price, one conversation.

---

## 2. CAC FOR FIRST 10 CUSTOMERS — IN DOLLARS

**$4,200 per customer acquired.**

Breakdown:
- **Founder outreach (cold + warm intro):** 5 hours per customer at $200/hr equivalent = $1,000
  - We're not hiring a sales rep. The CEO or engineering lead reaches out to targets they know or can find in tight circles (Slack communities, GitHub discussions, engineering Twitter).
- **Product customization/onboarding:** 2 hours per customer (setup, first run, calibration) = $400
- **Trial support (email + async debugging):** 3 hours per customer over 2-week trial = $600
- **Infrastructure cost (logs storage, compute for diffs):** ~$200 per customer during trial
- **Unallocated overhead (10% of founder time, office, tooling):** ~$2,000 spread across 10 = $200 per customer

**Total: $4,200 CAC to land the first 10.**

This assumes:
- We're not running ads.
- We're not hiring SDRs.
- Founder time is the lever. If this can't be sold by founder energy, it's not a fit for pre-revenue.

**Payback period:** $299 × 12 = $3,588 annual LTV (Year 1). Payback = 17 months. **This is unacceptable.**

You'll see why this matters in the final number.

---

## 3. RUNWAY COST TO VALIDATE THIS IDEA

**4 weeks of runway. ~0.6 months of our 18-month cushion.**

Assumes:
- 2-week build cycle (CLI tool, basic diff logic, log parsing, simple UI for output comparison).
- 2-week customer validation loop (identify 15 targets, reach out, get 3–5 into trial, measure engagement).
- No external costs (we use existing infra).
- 1 engineer (90% allocation), 1 co-founder (20% on outreach/support).

If after 4 weeks we don't have 3 teams in active trial or 1 qualified inbound lead expressing intent, we kill it and move on.

---

## 4. THE SINGLE NUMBER THAT HAS TO BE TRUE

**At least 2 out of the first 5 trial customers have to sign a contract within 30 days of starting the trial.**

This is the gate.

Here's why:
- A 40% conversion rate from trial → paid is the floor for "this is real demand, not polite interest."
- If only 1 in 5 converts, the CAC math becomes $21,000+ (assuming 10 customers total), and payback is 58+ months. We don't have that patience, and neither does the market.
- If 2 in 5 convert, CAC drops to ~$5,000, and payback is ~20 months. Still slow, but survivable if we can get volume.
- If 3 in 5 convert, we're at $3,500 CAC and 14-month payback—that's a business.

**How we measure it:**
- Signed contract = email with an actual payment method, or a Stripe subscription initiated. Not a verbal handshake, not "we're interested."
- 30 days = from first day of trial access to payment received.

---

## REAL TALK — WHY I'M NOT BLOCKING THIS, BUT NOT ENTHUSIASTIC

**The problem is real.** Flaky tests are a genuine tax. I believe teams will use this if it works.

**The financials are not encouraging yet:**
- A 17-month payback at $299/month means we need either (a) much higher price, or (b) much lower CAC, or (c) very high retention and expansion revenue.
- We can't do (b) at pre-revenue without brand; we've never sold anything.
- We *could* test (a)—pricing at $599 or $799—but that requires confidence that teams value time saved at that magnitude. We don't have that yet.
- (c) requires us to own the customer relationship and add features later. That's fine, but it's a longer play.

**What I'd rather see before we commit a full cycle:**
1. Spend 3 days (not a cycle) talking to 10 engineers at companies running 500+ tests. Ask: "If a tool gave you 20 minutes back per flaky investigation, what would you pay?" Record the range. If you hear $500+ from 5+ people, we run this cycle. If you hear $200 or less, we repriced or we kill it.
2. Build a **prototype in 1 week**—not production-ready, just functional enough to demo. A CLI that takes two test output files and diffs them side-by-side. Use it on one real codebase.
3. If the prototype feels like it actually saves time (not just theoretically), *then* we commit a 4-week cycle.

This is 1.5 weeks of founder energy. Very low cost. De-risks the pricing assumption.

**Alternative:** If you want to run the full 4-week cycle without validation, I'll support it—but I want the gate set clearly: 2 out of 5 first trials convert to paid, or we move on.

What's your read on whether we can talk to target engineers this week?
