package parser

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

// ASSUMPTION: JUnit XML is the primary format (common in CI/CD).
// Raw text format is a fallback for custom test runners.

// TestResult represents a single test outcome.
type TestResult struct {
	Name      string // test name
	Status    string // "pass", "fail", "error", "skipped"
	Output    string // full test output / assertion message
	Duration  float64
	Classname string // test class/package
}

// JUnitTestSuite is minimal JUnit XML unmarshaling.
type JUnitTestSuite struct {
	XMLName string      `xml:"testsuite"`
	Tests   []JUnitTest `xml:"testcase"`
}

// JUnitTest represents a single JUnit testcase element.
type JUnitTest struct {
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	Time      string `xml:"time,attr"`
	Failure   *struct {
		Message string `xml:"message,attr"`
		Text    string `xml:",chardata"`
	} `xml:"failure"`
	Error *struct {
		Message string `xml:"message,attr"`
		Text    string `xml:",chardata"`
	} `xml:"error"`
	Skipped *struct {
		Message string `xml:"message,attr"`
	} `xml:"skipped"`
}

// ParseJUnitXML reads a JUnit XML file and returns TestResult slice.
func ParseJUnitXML(path string) ([]TestResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var suite JUnitTestSuite
	if err := xml.Unmarshal(data, &suite); err != nil {
		return nil, fmt.Errorf("unmarshal XML: %w", err)
	}

	var results []TestResult
	for _, t := range suite.Tests {
		tr := TestResult{
			Name:      t.Name,
			Classname: t.Classname,
		}

		if t.Failure != nil {
			tr.Status = "fail"
			tr.Output = t.Failure.Message + "\n" + t.Failure.Text
		} else if t.Error != nil {
			tr.Status = "error"
			tr.Output = t.Error.Message + "\n" + t.Error.Text
		} else if t.Skipped != nil {
			tr.Status = "skipped"
			tr.Output = t.Skipped.Message
		} else {
			tr.Status = "pass"
			tr.Output = ""
		}

		results = append(results, tr)
	}

	return results, nil
}

// ParseRawText parses a simple text format: "TESTNAME:STATUS:OUTPUT".
// ASSUMPTION: simple fallback for custom test runners. Not production-grade.
func ParseRawText(path string) ([]TestResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	var results []TestResult
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 2 {
			continue
		}

		tr := TestResult{
			Name:   parts[0],
			Status: parts[1],
		}
		if len(parts) > 2 {
			tr.Output = parts[2]
		}
		results = append(results, tr)
	}

	return results, scanner.Err()
}