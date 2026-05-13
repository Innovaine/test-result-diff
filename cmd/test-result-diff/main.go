package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Innovaine/test-result-diff/internal/differ"
	"github.com/Innovaine/test-result-diff/internal/parser"
)

func main() {
	// ASSUMPTION: week 2 accepts file path arguments on CLI, not stdin. Simpler to test.
	baselinePath := flag.String("baseline", "", "Path to baseline test output file")
	currentPath := flag.String("current", "", "Path to current test output file")
	format := flag.String("format", "junit", "Test output format (junit, raw)")
	verbose := flag.Bool("v", false, "Verbose diff output")

	flag.Parse()

	if *baselinePath == "" || *currentPath == "" {
		fmt.Fprintf(os.Stderr, "Usage: test-result-diff -baseline <file> -current <file> [-format junit|raw] [-v]\n")
		os.Exit(1)
	}

	// Validate files exist.
	if _, err := os.Stat(*baselinePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: baseline file not found: %s\n", *baselinePath)
		os.Exit(1)
	}
	if _, err := os.Stat(*currentPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: current file not found: %s\n", *currentPath)
		os.Exit(1)
	}

	// Parse both outputs.
	var baselineTests, currentTests []parser.TestResult
	var parseErr error

	switch *format {
	case "junit":
		baselineTests, parseErr = parser.ParseJUnitXML(*baselinePath)
	case "raw":
		baselineTests, parseErr = parser.ParseRawText(*baselinePath)
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown format %s\n", *format)
		os.Exit(1)
	}

	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Error parsing baseline: %v\n", parseErr)
		os.Exit(1)
	}

	switch *format {
	case "junit":
		currentTests, parseErr = parser.ParseJUnitXML(*currentPath)
	case "raw":
		currentTests, parseErr = parser.ParseRawText(*currentPath)
	}

	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Error parsing current: %v\n", parseErr)
		os.Exit(1)
	}

	// Diff and output.
	diffs := differ.DiffTestResults(baselineTests, currentTests)

	if len(diffs) == 0 {
		fmt.Println("No differences found between baseline and current.")
		os.Exit(0)
	}

	// Print diff summary.
	fmt.Printf("Found %d flaky tests:\n\n", len(diffs))
	for _, d := range diffs {
		fmt.Printf("Test: %s\n", d.TestName)
		fmt.Printf("Status: %s → %s\n", d.BaselineStatus, d.CurrentStatus)
		if d.BaselineOutput != "" && d.CurrentOutput != "" {
			fmt.Println("Output Diff:")
			lineDiffs := differ.DiffLines(d.BaselineOutput, d.CurrentOutput)
			for _, ld := range lineDiffs {
				fmt.Printf("  %s\n", ld)
			}
		}
		if *verbose {
			fmt.Printf("Assertion Delta: %s\n", d.AssertionDelta)
		}
		fmt.Println()
	}
}