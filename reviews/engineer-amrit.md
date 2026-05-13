# Engineer review: "Test Result Diff — Automated test output comparison for flaky test diagnosis"

- **By:** Amrit (Engineer)
- **Cycle:** 1
- **Saved:** 13/05/2026, 5:19:15 AM

---

# Real Talk on Test Result Diff

I'll give you the honest build picture, not the pitch picture.

## 1. First usable version: 3 weeks, not 2

**Week 1:** CLI skeleton + file I/O (reading two test outputs side-by-side). Assertion parser that works for the top 3 test frameworks (pytest, Jest, Go testing). Basic diff logic—character-level, no intelligence yet. You ship *something* that runs and produces output.

**Week 2:** Normalize output across frameworks (same test can produce wildly different formats depending on the assertion library and verbosity level). Add CI log pull support (Jenkins, GitHub Actions, CircleCI)—this is the wall I'm about to describe. Build the rendering layer so diff output is actually *readable* in a terminal (color, alignment, truncation for large diffs).

**Week 3:** Edge case grinding. Timeouts. Malformed logs. Very large test runs (10k+ tests). Polish CLI interface. Basic documentation. Package it so a user can `pip install` or `npm install` and actually use it without a PhD.

The 3-week version is usable. The 2-week version is a proof-of-concept that looks done in a meeting but breaks on real CI output.

## 2. Where you hit the wall: CI log normalization

This is the silent killer in this idea.

Test output is a *nightmare* of inconsistency:
- **pytest** with `-v` flag produces one format; pytest with `--tb=short` produces another. Add plugins and you get custom output formats.
- **Jest** JSON reporter vs. human-readable output are two totally different shapes.
- **GitHub Actions** logs include timestamps, runner info, and ANSI color codes that need stripping.
- **Jenkins** logs can be plain text or wrapped in XML. CircleCI can be JSON or structured logs.
- A single test run can span multiple workers, and the order of output is not deterministic.

The pitch says "point to CI logs." That sounds simple. It's not. You need to:
1. Fetch the log (API keys, auth, per-provider code)
2. Parse it correctly (detect which framework ran, which runner, which format)
3. Reconstruct which test output belongs to which test (the test name might appear 3 times in a CI log)
4. Extract only the assertion failures and stack traces (not the setup noise)

This is a 1-week sink minimum. Maybe 1.5 weeks if you support 3+ CI providers and 3+ test frameworks.

If you underestimate this, you ship a tool that works beautifully on hand-pasted pytest output and then a user tries it with Jenkins + Go testing and it produces garbage. They uninstall in 10 minutes.

## 3. What I'd cut from the spec

**Cut v1:**
- Multi-provider CI log fetching (Week 1 only supports copy-paste test output, or a single CI provider you pick). Users can manually export logs until adoption justifies the API work.
- Flaky test trend tracking (tempting feature, zero revenue in it). Don't collect historical data. Don't show "this test has flaked 5 times this month." Just diff *this* pair of runs.
- Configuration files (`.testdiffrc`, custom assertion patterns, etc.). Sensible defaults or bust. Config is scope poison when you're pre-revenue.
- Web UI. Stay CLI-only. A web interface doubles the build time (backend, frontend, deployment, hosting).

**Keep v1:**
- Core diff logic: read two test outputs, find the assertions that changed, surface them clearly.
- Support for pytest and Jest (covers ~70% of the teams we'd target).
- Copy-paste or local file input only.
- Human-readable, color-coded diff output in the terminal.

## 4. Libraries that save a month

**Difflib (Python) or Diff-Match-Patch (any language):** Don't write diff logic from scratch. `difflib.unified_diff()` or `google-diff-match-patch` handles the heavy lifting. You just normalize the inputs and feed them in. Saves 3-4 days.

**Pytest's built-in JSON output plugin:** If you're targeting pytest, leverage `pytest --json-report`. Parsing JSON is trivial compared to parsing human-readable output. Jest has `--json` too. Default to structured output when available; fall back to regex parsing for edge cases.

**Click (Python) or Yargs (Node):** CLI framework. Don't hand-roll argument parsing. Click gives you help text, subcommands, and input validation for free. 1-2 days saved.

**Regex library for assertion extraction:** A single, well-tested library for finding assertion failures in various formats. Something like `regex` (Python) over the standard `re` module. Buy precision to avoid 10 edge cases later.

**Python:** Ship it in Python first. The CLI market is Python-native. Porting to Node or Go later is trivial once the logic is proven. You move faster in Python on a parsing and text-processing tool like this.

---

## The hard call: Should we build this?

I think **yes, but with a 2-week cap and clear kill condition.**

Here's why: Flaky test debugging is *actually painful*. Engineers will actually try this. The copy-paste + basic diff version is useful on its own—it doesn't need CI integration to be valuable. You can ship Week 1 output in 10 days, get real feedback, and decide if multi-provider log support is worth the second half of the cycle.

The risk: You spend 3 weeks building and discover that:
- Engineers don't actually use a CLI tool for this (they'd rather just fix the test).
- Assertion parsing is messier than expected and you can't cover 60% of cases without 4 more weeks.
- Users want historical trend data so badly that a basic diff isn't enough.

**Kill condition:** If by end of Week 2 you can't get 3 engineers from outside the company to use it unprompted and say "yeah, this saves me time," stop. Don't extend. Move on.

**What I need to see:** Two engineers using this on real flaky tests in their own CI pipelines by day 10. Not "they said it's cool," but "they ran it twice unprompted." That's your signal to continue.

Reputation risk is real here—we said Test Result Router was shipping and we ghosted it. I'm not repeating that. This one ships in 3 weeks or we don't ship it.
