package differ

import (
	"strings"

	"github.com/Innovaine/test-result-diff/internal/parser"
)

// Diff represents a difference between two test runs for the same test.
type Diff struct {
	TestName           string
	BaselineStatus     string
	CurrentStatus      string
	BaselineAssertion  string
	CurrentAssertion   string
	AssertionDelta     string
	IsFlaky            bool
}

// ComputeDiff compares two test result sets and returns flaky test diffs.
func ComputeDiff(baseline, current parser.TestResults) []Diff {
	baseline.Normalize()
	current.Normalize()

	var diffs []Diff

	baselineTests := baseline.AllTestIDs()
	currentTests := current.AllTestIDs()

	// Find tests that exist in both runs
	for testID := range baselineTests {
		if !currentTests[testID] {
			continue
		}

		baselineCase := baseline.GetTestByID(testID)
		currentCase := current.GetTestByID(testID)

		if baselineCase == nil || currentCase == nil {
			continue
		}

		// Check if status changed (flaky indicator)
		if baselineCase.Status != currentCase.Status {
			delta := computeAssertionDelta(baselineCase.AssertError, currentCase.AssertError)
			diffs = append(diffs, Diff{
				TestName:          testID,
				BaselineStatus:    baselineCase.Status,
				CurrentStatus:     currentCase.Status,
				BaselineAssertion: baselineCase.AssertError,
				CurrentAssertion:  currentCase.AssertError,
				AssertionDelta:    delta,
				IsFlaky:           true,
			})
		}
	}

	return diffs
}

// computeAssertionDelta returns a human-readable diff of assertion outputs.
// For now, simple line-by-line comparison; can be extended with a real diff algorithm.
func computeAssertionDelta(baseline, current string) string {
	if baseline == current {
		return ""
	}

	baselineLines := strings.Split(strings.TrimSpace(baseline), "\n")
	currentLines := strings.Split(strings.TrimSpace(current), "\n")

	var delta strings.Builder

	// Simple line-by-line comparison
	maxLines := len(baselineLines)
	if len(currentLines) > maxLines {
		maxLines = len(currentLines)
	}

	for i := 0; i < maxLines; i++ {
		bLine := ""
		cLine := ""

		if i < len(baselineLines) {
			bLine = baselineLines[i]
		}
		if i < len(currentLines) {
			cLine = currentLines[i]
		}

		if bLine != cLine {
			delta.WriteString("- " + bLine + "\n")
			delta.WriteString("+ " + cLine + "\n")
		}
	}

	return delta.String()
}