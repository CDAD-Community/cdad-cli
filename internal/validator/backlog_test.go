package validator

import (
	"strings"
	"testing"
)

func TestValidateBacklogPasses(t *testing.T) {
	backlog := `### EPIC-01 — Sample
**Status:** Proposed

##### STORY-01.1 — Do a thing
- **Status:** Ready
`
	report, err := validateBacklog(strings.NewReader(backlog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected a passing report, got %+v", report)
	}
}

func TestValidateBacklogDuplicateIDs(t *testing.T) {
	backlog := `### EPIC-01 — First
**Status:** Proposed

##### STORY-01.1 — First story
- **Status:** Ready

### EPIC-01 — Duplicate
**Status:** Proposed

##### STORY-01.1 — Duplicate story
- **Status:** Ready
`
	report, err := validateBacklog(strings.NewReader(backlog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !report.Failed() {
		t.Fatalf("expected duplicate IDs to fail validation")
	}
	if len(report.DuplicateEpicIDs) != 1 || report.DuplicateEpicIDs[0] != "EPIC-01" {
		t.Errorf("expected duplicate EPIC-01, got %v", report.DuplicateEpicIDs)
	}
	if len(report.DuplicateStoryIDs) != 1 || report.DuplicateStoryIDs[0] != "STORY-01.1" {
		t.Errorf("expected duplicate STORY-01.1, got %v", report.DuplicateStoryIDs)
	}
}

func TestValidateBacklogBadStatusVocabulary(t *testing.T) {
	backlog := `### EPIC-01 — Sample
**Status:** Draft

##### STORY-01.1 — Do a thing
- **Status:** Someday
`
	report, err := validateBacklog(strings.NewReader(backlog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !report.Failed() {
		t.Fatalf("expected bad status vocabulary to fail validation")
	}
	if len(report.BadEpicStatuses) != 1 || report.BadEpicStatuses[0] != "EPIC-01: Draft" {
		t.Errorf("unexpected bad epic statuses: %v", report.BadEpicStatuses)
	}
	if len(report.BadStoryStatuses) != 1 || report.BadStoryStatuses[0] != "STORY-01.1: Someday" {
		t.Errorf("unexpected bad story statuses: %v", report.BadStoryStatuses)
	}
}

func TestValidateBacklogEmptyEpicIsWarningOnly(t *testing.T) {
	backlog := `### EPIC-01 — No stories yet
**Status:** Proposed

## Next Section
`
	report, err := validateBacklog(strings.NewReader(backlog))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Failed() {
		t.Fatalf("an empty Epic must not fail validation, got %+v", report)
	}
	if len(report.EmptyEpics) != 1 || report.EmptyEpics[0] != "EPIC-01" {
		t.Errorf("expected EPIC-01 flagged as empty, got %v", report.EmptyEpics)
	}
}

func TestValidateBacklogMissingFile(t *testing.T) {
	_, err := ValidateBacklog("does-not-exist.md")
	if err == nil {
		t.Fatal("expected an error for a missing backlog file")
	}
}
