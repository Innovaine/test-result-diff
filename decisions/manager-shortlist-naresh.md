# Manager shortlist (cycle 1)

- **By:** Naresh (Manager)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:21:04 AM

---

I'm picking the two strongest candidates. Not because I love them—because they're the least bad, and we need to stop sitting on our hands.

**PICK 1: Test Result Router — Automatic CI test failure triage and owner assignment**

**Reason:** Carla nailed the buyer (single Engineering Manager with quarterly budget authority), Hassan gave us an honest 3-week build with real technical risk mitigation (the Slack queue problem is solvable), and Lucia locked in the floor gate: 2 paying customers by month-end or we kill it. Tomás is right that routing doesn't *fix* the problem, but Carla proved there's a buyer who feels 4-6 hours/week bleed and has $500-2K/quarter discretionary budget. That's a real lever. The per-run pricing is garbage (Lucia agrees), so we flip to $150/month flat and validate. Yes, it's a comfort problem masquerading as pain—but it's *their* comfort, and they're willing to write a check to reclaim it. That's the signal that matters.

The risk: Expansion ceiling is real (Lucia's right), but we don't need expansion—we need validation. If we hit 10 teams at $150/month, that's $18K ARR and proof of concept. We'll know if the model works by month 2.

---

**PICK 2: Test Result Diff — Automated test output comparison for flaky test diagnosis**

**Reason:** Amrit and Priya both say 5 weeks, not 2, but they also both confirm the core logic (diff two outputs, surface what changed) is sound and ships in Week 1. Lucia's gate is brutal but fair: 2 of 5 first trials convert to paid within 30 days, or we move on. That's a harder bar than Router (more selling effort required), but if we hit it, we've validated that the pain is sharp enough to activate spending. Tomás is right that we're selling aspirin, not the cure—but aspirin at $299/month with 40%+ trial-to-paid conversion is a business. The copy-paste + pytest/Jest version is genuinely useful out of the box. We're not waiting on CI integration.

The risk: Adoption friction is real (CLI tool, not web UI), and the payback is long (17–20 months). But that's not a blocker for pre-revenue validation—it's a metric we measure. If teams actually use it unprompted after Week 2, the model holds.

---

**Why not Test Result Router alone?**

Router is the safer pick (clearer buyer, faster sales motion), but it's also narrower—it only works for teams bleeding time on triage, and once adopted, the improvement is invisible (time gets reabsorbed into other work). Diff is riskier (CLI friction, adoption uncertainty), but if it lands, it's a tool people *actively use* multiple times per week, which drives retention and word-of-mouth.

I'm picking both because:
- **Router validates the buyer segment** (Engineering Managers at mid-market SaaS spend on DevTools).
- **Diff validates whether that segment will pay for *visible* time savings** (not just administrative efficiency).

If Router hits 2 customers by month-end and Diff can't get 2 of 5 trials to convert within 30 days, we know the segment buys infrastructure fixes, not symptom management. We pivot the portfolio accordingly.

If Diff converts 40%+ but Router stalls, we know the bottleneck is routing (buyer problem, not market problem), and we ship Diff first while fixing Router's GTM.

**Non-negotiable gates for both:**

**Router:**
- Carla's exact validation test by EOW: 5 calls to Engineering Managers, one question, record who says "4+ hours/week, I have budget now."
- Hassan's three blockers answered: (1) Do we have the backend engineer? (2) Are 3 customers lined up for day 8 integration? (3) Will you own weekly feedback calls?
- If we don't have clear answers by Monday, we slip the start date—don't guess.

**Diff:**
- Amrit's kill condition: 2 engineers from outside the company using it unprompted by end of Week 2 and saying it saves time. Not "that's cool"—actually running it on real flaky tests.
- Lucia's paywall: 2 of 5 first trials sign within 30 days at $299/month, or we stop.
- Priya's scope: GitHub Actions + pytest/Jest only. No web UI, no multi-framework, no historical trends. Cut everything else.

We're spending 18 weeks of runway combined across both—about 1 week of company burn. If both fail, we learn fast and have 17 months left to find something that works.

**Ship both or ship neither. No half-measures. I'm flagging that now.**
