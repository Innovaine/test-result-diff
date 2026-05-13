package differ

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Innovaine/test-result-diff/internal/parser"
)

// DiffResult represents the diff between baseline and current test.
type DiffResult struct {
	TestName         string
	BaselineStatus   string
	CurrentStatus    string
	BaselineOutput   string
	CurrentOutput    string
	AssertionDelta   string // regex-extracted assertion delta
}

// DiffTestResults compares baseline and current test runs, returns flaky tests.
// ASSUMPTION: a test is "flaky" if status changed OR output diverged.
func DiffTestResults(baseline, current []parser.TestResult) []DiffResult {
	// Index baseline by test name.
	baselineMap := make(map[string]parser.TestResult)
	for _, t := range baseline {
		baselineMap[t.Name] = t
	}

	var results []DiffResult

	for _, curr := range current {
		base, found := baselineMap[curr.Name]
		if !found {
			// Test didn't exist in baseline (new test). Skip for now.
			// ASSUMPTION: week 2 doesn't flag new tests as flaky; only status/output changes.
			continue
		}

		if base.Status != curr.Status || base.Output != curr.Output {
			// This test flaked.
			delta := extractAssertionDelta(base.Output, curr.Output)
			results = append(results, DiffResult{
				TestName:       curr.Name,
				BaselineStatus: base.Status,
				CurrentStatus:  curr.Status,
				BaselineOutput: base.Output,
				CurrentOutput:  curr.Output,
				AssertionDelta: delta,
			})
		}
	}

	return results
}

// DiffLines does a line-by-line diff between two text outputs.
// Returns lines prefixed with +/- to show changes.
// ASSUMPTION: simple diff; not a full Myers algorithm (can add later if needed).
func DiffLines(baseline, current string) []string {
	baseLines := strings.Split(baseline, "\n")
	currLines := strings.Split(current, "\n")

	var diffs []string

	// For now: naive approach. If baseline line != current, mark it.
	// TODO: Smith-Waterman or Myers diff if we need to handle large outputs.
	maxLen := len(baseLines)
	if len(currLines) > maxLen {
		maxLen = len(currLines)
	}

	for i := 0; i < maxLen; i++ {
		var baseLine, currLine string
		if i < len(baseLines) {
			baseLine = baseLines[i]
		}
		if i < len(currLines) {
			currLine = currLines[i]
		}

		if baseLine != currLine {
			if baseLine != "" {
				diffs = append(diffs, fmt.Sprintf("- %s", baseLine))
			}
			if currLine != "" {
				diffs = append(diffs, fmt.Sprintf("+ %s", currLine))
			}
		}
	}

	return diffs
}

// extractAssertionDelta uses regex to find assertion statements and diff them.
// ASSUMPTION: common assertion patterns (assertEquals, assert_equal, expect).
// Outputs the first differing assertion for readability.
func extractAssertionDelta(baseline, current string) string {
	// Pattern: match lines that look like assertions.
	// This is a strawman; can refine for specific test frameworks later.
	assertPattern := regexp.MustCompile(`(?i)(assert|expect|should).*`)

	baselineAssertions := assertPattern.FindAllString(baseline, -1)
	currentAssertions := assertPattern.FindAllString(current, -1)

	// Find first mismatch.
	for i := 0; i < len(baselineAssertions) && i < len(currentAssertions); i++ {
		if baselineAssertions[i] != currentAssertions[i] {
			return fmt.Sprintf("Expected: %s\nActual: %s", baselineAssertions[i], currentAssertions[i])
		}
	}

	// If no assertion mismatch found, return generic delta.
	if len(baselineAssertions) != len(currentAssertions) {
		return fmt.Sprintf("Assertion count changed: %d → %d", len(baselineAssertions), len(currentAssertions))
	}

	return "(no assertion delta extracted)"
}