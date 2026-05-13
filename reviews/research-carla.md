# Research review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Carla (Researcher)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:17:02 AM

---

# WHO PAYS

**Senior Engineer or Staff Engineer** (not manager) at mid-market SaaS: $50M–$500M ARR, 80–200 engineers, 1000+ tests/deploy cycle. Pull from **engineering budget line**, maybe $3–5K/year discretionary, maybe a $15–30K/year "DevTools" pool if they have one.

**NOT** the engineering manager. NOT the VP Eng. Staff engineers own flaky test triage *in practice* and feel the tax directly. They have credibility to ask for $200/month tooling. They also hate meetings about it—they just want the problem gone.

Why not smaller teams? Under 500 tests, flaky tests are an occasional annoyance, not a weekly time sink. The pain isn't sharp enough to spend money. Why not larger (1000+ test suites)? They've already bought a monitoring/analytics vendor (Buildkite, LaunchDarkly, custom dashboards) or buried the problem in CI configs. The buyer moves upmarket and the selling motion breaks.

---

# WHY THEY PAY NOW

A deploy cycle takes 20 minutes. A flaky test fails 1 in 8 runs. It takes the on-call engineer (usually a Senior or Staff person) 15–25 minutes to:
- Re-run locally.
- Check git blame.
- Read the test code.
- Read the assertion.
- Compare the output in their head or in a text editor.
- Conclude "environment timing issue" or "state leakage."
- Retry or skip.

Happens 2–3 times a week per engineer. 5–10 engineers on the team. That's **10–15 hours of senior engineering time weekly sunk into "why did this test flip?"**

At $150/hour loaded cost, that's $1500–2250/week in sunk labor. $200/month for a tool that collapses that to 3 minutes? Instant ROI.

**BUT:** This only registers as a budget line if the pain is *chronic*. A team that flakes once a month doesn't feel it. A team that flakes 3x a week is desperate.

---

# CLOSEST COMPETITOR / CURRENT STATE

**Doing nothing in a spreadsheet / grep / human eye:**
- Copy-paste test output into a text file.
- Open two terminal windows side by side.
- Grep for "AssertionError" or "expected vs actual."
- Diff them manually in VS Code.
- Still takes 10–15 minutes.

**Existing vendors (if any):**
- **Buildkite Insights, LaunchDarkly, Datadog CI monitoring:** Expensive ($500+/month), overkill for flaky diagnosis, require ingestion of full test telemetry. Not installed yet for most mid-market teams. Also require infrastructure buy-in.
- **Custom internal scripts:** Mature teams have written a bash script or small Python tool that diffs JSON test reports. It's brittle, unmaintained, team-specific.
- **LLM-based "AI debugging":** Anthropic Claude, OpenAI, etc. Unproven for this use case. Probably overkill. Might work but no clear shipping path.

**Your actual competitor:** *The 15 minutes they spend right now doing it by hand.* Not another SaaS vendor. Internal labor.

---

# THE ONE NUMBER

For $1M ARR:

- **Market:** 10,000 teams running 500+ tests/cycle in mid-market SaaS (conservative—likely 50K+, but be pessimistic).
- **Adoption:** 5% of addressable market converts and stays (500 teams).
- **Price:** $200/month = $2,400/year.
- **ARR:** 500 × $2,400 = **$1.2M.**

**Sensitivity:**
- If adoption is 2% (200 teams): ARR = $480K. Survive, no growth.
- If adoption is 10% (1,000 teams): ARR = $2.4M. Good.
- If price is $100/month (to compete on cost): ARR = $600K at 5% adoption. Below threshold.
- If price is $400/month (premium): ARR = $2.4M at 5% adoption. Tight but workable.

**The kill-line:** Adoption below 2% or price below $150/month makes this sub-$1M on addressable market of 10K teams. You need **either a bigger market (are there 50K+ such teams?) or higher adoption (can you hit 10%+?).**

---

# THIS WEEK'S VALIDATION TEST

**Who to talk to:**
1. Staff Engineer at Stripe, Square, Twilio (mid-market SaaS, known to have heavy CI/CD cultures).
2. Staff Engineer at Figma, Notion, etc. (unicorn-stage, likely have flaky test pain).
3. Engineering lead at a $50–200M ARR SaaS you can find via LinkedIn (search "Staff Engineer" + "CI/CD" or "test infrastructure").
4. (If you have a warm intro:) The on-call engineer who runs CI at a Sequoia/Andreessen portfolio company.
5. Your own network: someone who works in DevOps or test infrastructure at a company you trust to give honest feedback.

**The one question:**
> "When a test flakes—fails once, then passes again without code changes—how do you figure out what changed? Walk me through the last time it happened."

**Listen for:**
- Do they describe a 5-minute process or a 20-minute process?
- Do they pull test output diffs manually or do they have tooling?
- Do they say "we just re-run" or "we debug the hell out of it"?
- Does the pain register as annoying or as time-stealing?

**What kills the idea:**
- Answer: "We don't get flaky tests." (Means: either tiny test suite, or they've already solved it, or they don't measure.)
- Answer: "We use [Buildkite Insights / Datadog / custom platform] and it's fine." (Means: buyer already has solution; you're a me-too.)
- Answer: "Takes maybe 5 minutes, not a big deal." (Means: pain threshold too low; adoption rate tanks.)
- Answer: "We'd need this integrated into [Slack / GitHub / our CI platform]." (Means: CLI tool doesn't fit workflow; you need to rebuild as a plugin—2x engineering lift.)

**What keeps it alive:**
- Answer: "Yeah, I spend 15–20 minutes figuring out what flipped. It's the worst part of my week." (Bingo. Real, quantified pain.)
- Answer: "I write a script every time. Would pay for something that just works." (Means: market exists, willingness to pay is real.)
- Follow-up if answer is positive: "Would $200/month be worth it if this cut that 20 minutes to 3?" (Get explicit price acceptance before you build.)

**When:** This week. 3 calls minimum. 30 minutes each. Report back with exact quotes and one clear "this person will buy" or "this person killed it" signal.

---

## My take (you need to hear this):

This is better than Test Result Router. The buyer is clearer (Staff Engineer, not "DevOps lead"). The pain is sharper (time cost, not someone else's problem). The competitor is simpler (15 minutes of manual work, not a vendor category).

But you're still pre-revenue on a 2-week cycle. Do NOT build a full product. Build the minimum:
- A shell script or Python CLI that diffs two test JSON outputs.
- Takes 3 hours to code, 1 day to test.
- By end of week, you know if the buyer exists.

If all three calls say "yes, I'd use that"—you have validation for a 2-week sprint. If all three say "we're fine" or "it's too integrated to our platform"—you kill it and move to the next one.

Ship the question first. Ship the code second.
