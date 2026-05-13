# Test Result Diff

Automated test output comparison for flaky test diagnosis.

## What it does

Compares test output between two runs (baseline and current), identifies flaky tests, and generates structured diffs of assertion output deltas.

**Current scope:**
- CLI binary (no UI, no integrations)
- JUnit XML parsing
- Line-by-line diff + regex-based assertion extraction
- Plain text output (GitHub Action integration coming next week)

## Build
