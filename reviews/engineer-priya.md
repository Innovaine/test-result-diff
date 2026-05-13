# Engineer review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Priya (Engineer)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:20:11 AM

---

I'm going to be direct because that's how I work, and because this pitch has the same shape as the Router we just reviewed — underestimated scope, overestimated parsing complexity, and a deployment assumption that doesn't survive contact with real CI systems.

---

## 1. How long the first usable version really takes

**5 weeks, not 2.**

I'm splitting that into:
- **Week 1-2: Parse + diff logic** — You need to handle GitHub Actions, CircleCI, and Jenkins output formats at minimum. Each has its own timestamp style, assertion formatting, variable names. Not hard, but every edge case costs a day. You'll hit "wait, do we diff raw output or AST" questions around day 4.
- **Week 2-3: CLI + local file input** — Takes 3-4 days to build something that doesn't break on whitespace or encoding issues. Testing on real CI output adds another week because you'll discover your regex doesn't handle multiline stack traces or colored terminal output until you point it at actual logs.
- **Week 3-4: The "wait, people actually want this" wall** — You'll ship a 0.1 and immediately learn that "paste two outputs" doesn't work for anyone; they want it to point at CI job IDs or URLs. That's another 5-7 days of plumbing.
- **Week 5: Polish, docs, error handling** — Nobody installs a tool that crashes on bad input or has no help text.

The "30 seconds" in your pitch assumes the inputs are already clean and formatted the way you expect. They won't be.

---

## 2. Where you'd hit a wall

Three hard stops:

**Wall 1: Test output format fragmentation**  
GitHub Actions, CircleCI, Jenkins, GitLab CI, Buildkite — they all log test results differently. `pytest` output looks nothing like `Jest` output. You can't build a general-purpose diff tool without deciding: do you support all formats day one, or do you pick one and accept that 40% of your TAM can't use it?

I'd pick one (GitHub Actions + pytest, since it's most common in mid-market engineering teams), ship that, and iterate. But you have to make that call *before* you start writing the parser, not halfway through.

**Wall 2: Context matching**  
"This assertion changed between runs" is useless if you can't tell the engineer which test case produced it. You need to match test names across runs, handle parameterized tests, and deal with tests that don't have stable identifiers. That's not a parsing problem — that's a deduplication + heuristic matching problem. Expect 4-5 days of debugging just on "why is test case #3 matching to two different runs?"

**Wall 3: Deployment assumption**  
Your pitch says "install it in one standup." That assumes:
- Teams have a shared machine or Slack integration to run CLI tools
- Or they're comfortable installing a Python/Node binary locally
- Or they use it in their CI itself

Reality: most teams want a web UI. They'll ask for it day one. You don't build that in week 1. You stub it and accept that adoption is slower than your pitch promises, or you build it and blow the timeline.

---

## 3. What I'd cut from the spec

**Cut immediately:**
- Web UI, cloud hosting, auth, accounts — ship CLI only. No database.
- Support for more than two test frameworks on day one. GitHub Actions + pytest. That's it.
- Integration with CI systems as a plugin. No — just take log files as input, or URLs that the user downloads first.
- "Email alerts" or "Slack notifications" — not in v1.
- Performance optimization for 10,000-test suites — optimize for the 500-test case, the one that actually hurts.

**Keep:**
- Side-by-side diff output that's readable in terminal
- "Show me which assertions actually changed" — that's the 80/20
- Error messages that don't suck (because a bad error message kills adoption faster than missing features)
- A single, stable CLI interface — don't rip it out in week 3

---

## 4. Which existing tools/libraries save a month

**difflib (Python stdlib)**  
Use `difflib.unified_diff()` or `ndiff()` for the comparison engine. Don't build your own. You'll spend three days on whitespace handling and lose anyway.

**pytest's JSON plugin output**  
If you target pytest specifically, force users to run with `--json-report`. It's already structured, already stable, and you skip 80% of the parsing work. Yes, it's a constraint. It's also a month saved.

**Click (Python CLI library)**  
Build the CLI with Click, not argparse. It handles help text, subcommands, and error messaging. You'll need it to be bulletproof because a confusing CLI kills a tool faster than a missing feature.

**colorama (for cross-platform terminal color)**  
Terminal output needs to be color-coded (failed assertions in red, passed in green, changed in yellow). Use colorama so it doesn't break on Windows. One afternoon's work, not two weeks of debugging terminal weirdness.

**gitpython or similar for git blame (if you go there)**  
If you eventually want to connect a changed assertion to the commit that broke it, you'll need git history. Don't do this in v1. But when you do, gitpython saves the integration work.

**Don't use:**
- Heavy test frameworks (Cypress, Playwright, etc.) — you're not testing the tool, you're testing *other people's* test output.
- Elasticsearch or any real database — you're processing two files per invocation. A database adds complexity and zero value.

---

## My call on this one

**Ship it. But reset expectations.**

This is a real problem. Flaky test diagnosis is genuinely a 20-minute tax per incident, and if you solve it for one framework, you have a beachhead. But:

1. **Start with CLI-only, GitHub Actions + pytest, local file input.** No web UI, no CI plugin, no multi-framework support. That gets you to "real users can use this" in 4 weeks.

2. **Week 4, not week 2, before anyone outside the team touches it.** Your pitch says "install in a standup" — that only works if the tool is boring and reliable. Boring takes time.

3. **Have a paying customer lined up before launch, or accept 6-month sales cycle.** You're right that teams will pay $200/month for the time saved. But "install at standup" adoption doesn't happen without word-of-mouth, and word-of-mouth needs someone who's already using it and loves it. That's a customer, not a free trial.

4. **Pick your battles on format support early.** If you say "pytest only," some teams walk. If you try to support everything, you ship late and broken. I'd take 60% TAM over 0% coverage.

I'm not opposed to this. I'm opposed to the timeline and the assumption that it'll be trivial to parse. The core insight (diff test output, find the delta) is sound. The execution is a 5-week sprint, not a 2-week hack.

What's your actual constraint here — do you have a team member ready to own the pytest parser, or is this speculative?
