// Package validator implements the deterministic, LLM-free governance
// checks behind `cdad validate` (see cdad/context/architecture.md,
// internal/validator: Domain only, never Network or LLM/agent calls).
//
// BacklogReport ports the exact rules already agreed and shipped as
// cdad/scripts/cdad-check-backlog.sh, so the CLI enforces the same
// contract natively (including on Windows, without bash) instead of a
// second, drifting definition of "valid backlog".
package validator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

var (
	epicHeaderRe   = regexp.MustCompile(`^### (EPIC-[0-9]+)`)
	storyHeaderRe  = regexp.MustCompile(`^##### (STORY-[0-9]+(?:\.[0-9]+)?)`)
	epicStatusRe   = regexp.MustCompile(`^\*\*Status:\*\*\s*(.+?)\s*$`)
	storyStatusRe  = regexp.MustCompile(`^\s*- \*\*Status:\*\*\s*(.+?)\s*$`)
	level2HeaderRe = regexp.MustCompile(`^## `)
)

// Allowed status vocabulary, exactly as cdad-check-backlog.sh defines it.
var (
	epicStatuses = map[string]bool{
		"Proposed": true, "Planned": true, "In Progress": true,
		"Completed": true, "Cancelled": true,
	}
	storyStatuses = map[string]bool{
		"Proposed": true, "Ready": true, "In Progress": true,
		"Blocked": true, "Done": true, "Cancelled": true,
	}
)

// BacklogReport is the outcome of validating one backlog.md.
type BacklogReport struct {
	DuplicateEpicIDs  []string
	DuplicateStoryIDs []string
	BadEpicStatuses   []string // "EPIC-01: <offending status>"
	BadStoryStatuses  []string // "STORY-01.1: <offending status>"
	EmptyEpics        []string // warning only, not a failure
}

// Failed reports whether any structural-integrity rule was violated.
// EmptyEpics never fails validation: a freshly proposed Epic legitimately
// has no Stories yet.
func (r *BacklogReport) Failed() bool {
	return len(r.DuplicateEpicIDs) > 0 || len(r.DuplicateStoryIDs) > 0 ||
		len(r.BadEpicStatuses) > 0 || len(r.BadStoryStatuses) > 0
}

// ValidateBacklog reads path and checks Epic/Story ID uniqueness and
// status vocabulary. A missing file is reported via the os.IsNotExist
// error, mirroring the shell script's "cannot determine" exit code.
func ValidateBacklog(path string) (*BacklogReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return validateBacklog(f)
}

func validateBacklog(r io.Reader) (*BacklogReport, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	report := &BacklogReport{}
	epicCount := map[string]int{}
	storyCount := map[string]int{}

	var currentEpic string
	currentEpicStories := 0

	closeEpic := func() {
		if currentEpic != "" && currentEpicStories == 0 {
			report.EmptyEpics = append(report.EmptyEpics, currentEpic)
		}
		currentEpic = ""
		currentEpicStories = 0
	}

	for i, line := range lines {
		switch {
		case epicHeaderRe.MatchString(line):
			closeEpic()
			id := epicHeaderRe.FindStringSubmatch(line)[1]
			epicCount[id]++
			currentEpic = id
			if status, ok := lookahead(lines, i, epicStatusRe); ok && !epicStatuses[status] {
				report.BadEpicStatuses = append(report.BadEpicStatuses, fmt.Sprintf("%s: %s", id, status))
			}
		case storyHeaderRe.MatchString(line):
			id := storyHeaderRe.FindStringSubmatch(line)[1]
			storyCount[id]++
			currentEpicStories++
			if status, ok := lookahead(lines, i, storyStatusRe); ok && !storyStatuses[status] {
				report.BadStoryStatuses = append(report.BadStoryStatuses, fmt.Sprintf("%s: %s", id, status))
			}
		case level2HeaderRe.MatchString(line):
			closeEpic()
		}
	}
	closeEpic()

	report.DuplicateEpicIDs = duplicates(epicCount)
	report.DuplicateStoryIDs = duplicates(storyCount)

	return report, nil
}

// lookahead mirrors `grep -A2` against a header line: it checks the two
// lines immediately following i for a match against re.
func lookahead(lines []string, i int, re *regexp.Regexp) (string, bool) {
	for j := i + 1; j <= i+2 && j < len(lines); j++ {
		if m := re.FindStringSubmatch(lines[j]); m != nil {
			return m[1], true
		}
	}
	return "", false
}

func duplicates(counts map[string]int) []string {
	var out []string
	for id, n := range counts {
		if n > 1 {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}
