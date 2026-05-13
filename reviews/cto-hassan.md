# CTO review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Hassan (CTO)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:17:52 AM

---

# AGGRESSIVE SCOPE CUT — Test Result Diff

## 1. BUILD COMPLEXITY — 7-DAY UGLY MVP

**What ships:**
- CLI binary (Go) — takes two test output files or raw paste stdin
- Diffs JSON/XML test payloads line-by-line, highlights changed assertions
- Outputs colored terminal diff (green pass → red fail, show the delta)
- No parsing intelligence — assumes standard formats (pytest JSON, JUnit XML, Go test output). If it's not one of those three, user re-runs with `--format=raw` and pastes the plaintext
- Hardcoded to those three formats. If a fourth format shows up, user files an issue; we add it in cycle 2

**What does NOT ship:**
- Web UI (cut entirely — CLI only)
- CI integration / webhook listeners (cut — manual paste/file upload only)
- Database / result history (cut — stateless)
- Authentication, multi-team namespacing, RBAC (cut — single-user CLI tool)
- Test failure prediction / "flake scoring" (cut — diffs only, no ML)
- Slack notification integration (cut — user runs tool manually)

**Timeline:**
- Day 1–2: Go CLI skeleton, JSON/XML parser library (use existing `encoding/json` + `github.com/antchfx/xmlquery`), file input handler
- Day 3–4: Diff algorithm (simple line-by-line + assertion extraction), colored output formatting
- Day 5: pytest + JUnit + Go format handling, edge cases (malformed input, empty diffs)
- Day 6: CLI flag parsing, help text, README with copy-paste examples
- Day 7: Manual testing against real CI output from a volunteer team, fix crashes, ship

**Ugly truth:** If a test result doesn't parse cleanly, we show a raw diff. Not elegant. Fixes come later.

---

## 2. TEAM / SKILLS NEEDED

**You need:**
- **1 senior backend engineer** (4+ years Go experience, has shipped CLI tools before) — this person leads, owns architecture decisions, ships it themselves if needed. NOT a junior; the format parsing and diff logic need confidence. This is the main build thread.
- **1 mid-level engineer** (2+ years, any backend language, comfortable with Go) — handles test format integration, runs manual validation against real CI logs, writes the README and examples.
- **You (Hassan)** — 3 hours total (day 1 architecture review, day 5 risk check, day 7 sign-off). Not hands-on; judgment gates only.

**You do NOT need:**
- Frontend person (no UI)
- DevOps/infra person (it's a stateless binary)
- QA automation (manual testing with volunteer teams is enough for MVP)
- Product manager (the pitch is already the spec)

**Hiring reality:** If you don't have a senior Go person on staff, this scope dies. Don't hire for it in 7 days. If you don't have them, we pivot to Python + CLI instead, or we push the ship date to day 10 and add one. Call that now.

---

## 3. BIGGEST TECHNICAL RISK THAT KILLS US IN PRODUCTION

**The risk:** Test output parsing fails silently or crashes on edge cases (truncated logs, encoding mismatches, mixed format outputs). A user runs `test-diff log1.json log2.json`, gets no output or a crash, and deletes the tool.

**Why it kills us:** No fallback. No dashboard. No support. If the CLI crashes or hangs, the user's next action is `rm -rf` and "this thing is garbage."

**Mitigation (non-negotiable):**
1. **Panic handlers** — any parse failure catches panic, logs the raw input to stderr, suggests `--format=raw` as fallback, exits with code 1 (not a crash)
2. **Timeout on diff** — if a test result is >100MB or diff takes >5 seconds, bail with a message ("output too large, file a feature request")
3. **Validation before diff** — parse both files completely before attempting diff. If either parse fails, exit with "could not parse [filename], try `test-diff --debug [file]` to see what failed"
4. **Manual testing on real CI output** — day 6, grab 3 real test logs from a production team. If it crashes on one, we fix it before ship. If it doesn't crash, we ship.

**Testing:** No formal QA suite (too slow). Senior engineer manually runs it against the three formats + one deliberately broken log. If it doesn't crash, we ship.

---

## 4. THREE FEATURES TO CUT

### **CUT #1: Multi-run aggregation ("show me flakiness over 10 runs")**
- **Why it's tempting:** Teams want to see patterns, not just one-shot diffs. Looks "comprehensive."
- **Why we cut it:** Requires result storage, query logic, more CLI args. Pushes build to day 10+. MVP is "what changed between these two runs?" not "is this test actually flaky?"
- **Defer to:** Cycle 2, if paying customers ask for it.

### **CUT #2: Assertion-level metadata (show which assertion failed, suggest fix)**
- **Why it's tempting:** "Smart" — the tool could say "assertion.equals(5, actual) failed because actual=7" and suggest a fix direction.
- **Why we cut it:** Requires AST parsing or regex heuristics per language (Python assertions look different from Go). Fragile. We'd ship something that works 60% of the time and breaks silently.
- **Defer to:** Cycle 2 or 3. Start with raw delta; teams will tell us what assertions they care about most.

### **CUT #3: CI log pulling (direct GitHub Actions / Jenkins / GitLab integration)**
- **Why it's tempting:** "Seamless integration" — users don't have to leave their CI dashboard.
- **Why we cut it:** Requires API keys, OAuth, permission handling. Adds auth complexity we said NO to. Day 1 user copies-pastes two outputs or points to local files. That works.
- **Defer to:** Cycle 2 if adoption is strong. Then we do GitHub Actions integration as a paid add-on.

---

## HARD DECISIONS

**Ship date: Day 7, Tuesday EOD.** No slip. If we're not done, we cut the Go format handler and ship with pytest + JUnit only. Users can request Go support in the issue tracker.

**QA = three real test runs manually tested.** We don't have time for a test harness. The senior engineer is our QA.

**Pricing: $200/month.** We'll validate that later. For now, build a tool so obviously useful that one engineer shows it to their team lead, and the team decides to pay. No sales deck, no feature comparison. Just "try it for free for a week, then it's $200/month if you keep using it."

**Success metric for cycle 2 handoff:** One team (real, not friends-and-family) is using it regularly enough to consider paying. If we don't have that by end of cycle 1, we kill this product and ship something else. No extended validation runway.

---

## MY ACTUAL CONCERN

I'm not worried about the code. A senior Go engineer ships this in 7 days. I'm worried about **whether teams will actually use a CLI tool for this.**

The pitch assumes engineers love CLIs. Most don't — they want to stay in Slack or their CI dashboard. If we ship this and three months in realize "nobody's copy-pasting logs into a terminal," we've wasted effort.

**Validation I'd want before we ship:**
- Grab two volunteer teams (not friends, real engineering teams from YC network or Slack communities). Send them the early binary on day 5. Have them run it once against real logs. Do they immediately see value, or do they say "cool but I'm not doing this"?

If both say "cool but," we pivot the output format to "GitHub comments on failed CI runs" instead of "CLI" — same diff engine, different UX. That's doable by day 7 too, but a different build.

**What I need from you before we start:**
- Do we have a senior Go engineer (4+ years, shipped CLI) on staff or available this week? If no, we defer ship to day 10 or pivot to Python.
- Can you commit 3 hours for gates (day 1, 5, 7)? If not, this gets delayed — I don't want the senior engineer second-guessing architecture mid-build.
- Should we validate the "do engineers actually want a CLI for this" hypothesis before we commit, or do we ship ugly and learn?

---

**My vote:** Ship the CLI. It's testable, it's fast, it's honest about what it does. If it flops, we know in two weeks. That's better than four months of "maybe we should add web UI."
