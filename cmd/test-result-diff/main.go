package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Innovaine/test-result-diff/internal/differ"
	"github.com/Innovaine/test-result-diff/internal/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: test-result-diff <baseline.xml> <current.xml>\n")
		os.Exit(1)
	}

	baselineFile := args[0]
	currentFile := args[1]

	// Parse baseline results
	baselineData, err := os.ReadFile(baselineFile)
	if err != nil {
		log.Fatalf("Failed to read baseline file: %v", err)
	}

	var baselineResults parser.TestResults
	if err := xml.Unmarshal(baselineData, &baselineResults); err != nil {
		log.Fatalf("Failed to parse baseline XML: %v", err)
	}

	// Parse current results
	currentData, err := os.ReadFile(currentFile)
	if err != nil {
		log.Fatalf("Failed to read current file: %v", err)
	}

	var currentResults parser.TestResults
	if err := xml.Unmarshal(currentData, &currentResults); err != nil {
		log.Fatalf("Failed to parse current XML: %v", err)
	}

	// Compute diff
	diffs := differ.ComputeDiff(baselineResults, currentResults)

	// Output structured diff
	if len(diffs) == 0 {
		fmt.Println("No differences found.")
		return
	}

	fmt.Println("=== Test Result Diff ===")
	for _, diff := range diffs {
		fmt.Printf("\n[FLAKY] %s\n", diff.TestName)
		fmt.Printf("Status: %s → %s\n", diff.BaselineStatus, diff.CurrentStatus)
		if diff.BaselineAssertion != "" {
			fmt.Printf("Baseline assertion:\n%s\n", diff.BaselineAssertion)
		}
		if diff.CurrentAssertion != "" {
			fmt.Printf("Current assertion:\n%s\n", diff.CurrentAssertion)
		}
		if diff.AssertionDelta != "" {
			fmt.Printf("Delta:\n%s\n", diff.AssertionDelta)
		}
	}
}