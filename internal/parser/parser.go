package parser

import (
	"encoding/xml"
)

// TestResults represents the root of a JUnit-style test output XML.
type TestResults struct {
	XMLName  xml.Name  `xml:"testsuites"`
	Suites   []Suite   `xml:"testsuite"`
	Tests    int       `xml:"tests,attr"`
	Failures int       `xml:"failures,attr"`
	Skipped  int       `xml:"skipped,attr"`
}

// Suite represents a single test suite.
type Suite struct {
	XMLName   xml.Name `xml:"testsuite"`
	Name      string   `xml:"name,attr"`
	Tests     int      `xml:"tests,attr"`
	Failures  int      `xml:"failures,attr"`
	Skipped   int      `xml:"skipped,attr"`
	Time      string   `xml:"time,attr"`
	TestCases []Case   `xml:"testcase"`
}

// Case represents a single test case.
type Case struct {
	XMLName     xml.Name `xml:"testcase"`
	Name        string   `xml:"name,attr"`
	ClassName   string   `xml:"classname,attr"`
	Time        string   `xml:"time,attr"`
	Status      string   // Computed: "pass", "fail", "skip"
	Failure     *Failure `xml:"failure"`
	Skipped     *Skipped `xml:"skipped"`
	AssertError string   // Raw assertion/error output
}

// Failure represents a test failure.
type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Text    string   `xml:",chardata"`
}

// Skipped represents a skipped test.
type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr"`
}

// Normalize normalizes the parsed results: computes test status for each case.
func (tr *TestResults) Normalize() {
	for i := range tr.Suites {
		for j := range tr.Suites[i].TestCases {
			case_ := &tr.Suites[i].TestCases[j]
			if case_.Failure != nil {
				case_.Status = "fail"
				case_.AssertError = case_.Failure.Text
			} else if case_.Skipped != nil {
				case_.Status = "skip"
			} else {
				case_.Status = "pass"
			}
		}
	}
}

// GetTestByID returns a test case by its fully qualified name (ClassName.Name).
func (tr *TestResults) GetTestByID(id string) *Case {
	for i := range tr.Suites {
		for j := range tr.Suites[i].TestCases {
			case_ := &tr.Suites[i].TestCases[j]
			fqn := case_.ClassName + "." + case_.Name
			if fqn == id {
				return case_
			}
		}
	}
	return nil
}

// AllTestIDs returns a set of all test case FQNs in the result.
func (tr *TestResults) AllTestIDs() map[string]bool {
	ids := make(map[string]bool)
	for i := range tr.Suites {
		for j := range tr.Suites[i].TestCases {
			case_ := &tr.Suites[i].TestCases[j]
			fqn := case_.ClassName + "." + case_.Name
			ids[fqn] = true
		}
	}
	return ids
}