// Package bootstrap implements the deterministic, non-agent part of
// `cdad init`: detecting the target project and creating the mechanical
// CDAD structure the CLI itself owns.
//
// It deliberately stops short of the full CDAD workspace contract in
// AGENTS.md. The portable-core files (AGENTS.md, README-CDAD.md and its
// Spanish translation) and the generic cdad/ documents (INDEX.md,
// CHANGE-REQUEST.md, INSTALLATION*.md, USAGE*.md, the ADR template, and
// cdad/README.md) are canonical content that must ship identically to
// every project; how the CLI obtains and versions that content is not yet
// decided (see cdad/context/stack.md, internal/artifacts row). Inventing
// that content here would be an unratified architecture decision, so
// those paths are reported as pending instead of guessed at.
//
// cdad/context/ and cdad/adr/ are created empty: their content is always
// authored by the AI agent, never by this CLI (see
// cdad/context/architecture.md, "Data model ownership").
package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/CDAD-Community/cdad-cli/internal/domain"
)

// mechanicalDirs are created empty; their contents are authored later by
// either the AI agent (context/, adr/) or the project owner (docs/,
// proposals/, scripts/).
var mechanicalDirs = []string{
	filepath.Join("cdad", "adr"),
	filepath.Join("cdad", "context"),
	filepath.Join("cdad", "docs"),
	filepath.Join("cdad", "proposals"),
	filepath.Join("cdad", "scripts"),
}

// mechanicalFiles are process/audit scaffolding this CLI owns outright, so
// it may author their starting content itself.
var mechanicalFiles = map[string]string{
	filepath.Join("cdad", "backlog.md"): backlogStarter,
}

// pendingArtifactFiles are portable-core/generic CDAD documents this
// story does not create: sourcing their canonical content is a separate,
// not-yet-decided concern (bootstrap artifact acquisition).
var pendingArtifactFiles = []string{
	"AGENTS.md",
	"README-CDAD.md",
	"README-CDAD.es.md",
	filepath.Join("cdad", "README.md"),
	filepath.Join("cdad", "INDEX.md"),
	filepath.Join("cdad", "CHANGE-REQUEST.md"),
	filepath.Join("cdad", "INSTALLATION.md"),
	filepath.Join("cdad", "INSTALLATION.es.md"),
	filepath.Join("cdad", "USAGE.md"),
	filepath.Join("cdad", "USAGE.es.md"),
	filepath.Join("cdad", "adr", "ADR-TEMPLATE.md"),
}

const completionFileName = "CDAD-COMPLETION.md"

const backlogStarter = `# Project Backlog

> The project's living development line -- Epics, Stories, and the work
> currently expected to be built. This is a development-planning artifact,
> not architecture.

No Epics or Stories are defined yet.

## Current Focus

- None

## Next Work

- None

## Blocked

- None
`

const completionHeader = `# CDAD Bootstrap -- Completion Report

> The durable record of what CDAD initialization did in this project, and
> the outcome of any later re-run. Append, do not overwrite.

`

// Report is the outcome of one Init run.
type Report struct {
	Version   string
	Created   []string
	Conflicts []string
	Pending   []string
}

// Init detects root as an existing project directory and creates the
// mechanical CDAD structure inside it, never overwriting anything already
// present. version identifies the CLI build and is recorded in
// cdad/CDAD-COMPLETION.md.
func Init(project domain.CDADProject, version string) (*Report, error) {
	root := project.RootPath

	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("project directory not detected: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", root)
	}

	report := &Report{Version: version}

	for _, dir := range mechanicalDirs {
		full := filepath.Join(root, dir)
		if st, err := os.Stat(full); err == nil && st.IsDir() {
			report.Conflicts = append(report.Conflicts, dir+string(filepath.Separator))
			continue
		}
		if err := os.MkdirAll(full, 0o755); err != nil {
			return nil, fmt.Errorf("creating %s: %w", dir, err)
		}
		report.Created = append(report.Created, dir+string(filepath.Separator))
	}

	for path, content := range mechanicalFiles {
		full := filepath.Join(root, path)
		if _, err := os.Stat(full); err == nil {
			report.Conflicts = append(report.Conflicts, path)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			return nil, fmt.Errorf("creating parent of %s: %w", path, err)
		}
		if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
			return nil, fmt.Errorf("writing %s: %w", path, err)
		}
		report.Created = append(report.Created, path)
	}

	for _, path := range pendingArtifactFiles {
		full := filepath.Join(root, path)
		if _, err := os.Stat(full); err == nil {
			report.Conflicts = append(report.Conflicts, path)
			continue
		}
		report.Pending = append(report.Pending, path)
	}

	if err := recordCompletion(root, report); err != nil {
		return nil, fmt.Errorf("recording completion entry: %w", err)
	}

	return report, nil
}

// recordCompletion appends a dated entry to cdad/CDAD-COMPLETION.md,
// creating the file with its header first if it does not exist yet.
func recordCompletion(root string, report *Report) error {
	path := filepath.Join(root, "cdad", completionFileName)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(completionHeader), 0o644); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	entry := fmt.Sprintf(
		"## cdad init -- %s\n\nCLI version: %s\n\nCreated:\n%s\nConflicts (left untouched):\n%s\nPending (requires bootstrap artifact acquisition, not yet implemented):\n%s\n---\n\n",
		time.Now().UTC().Format("2006-01-02"),
		report.Version,
		bulletList(report.Created),
		bulletList(report.Conflicts),
		bulletList(report.Pending),
	)

	_, err = f.WriteString(entry)
	return err
}

func bulletList(items []string) string {
	if len(items) == 0 {
		return "- none\n"
	}
	out := ""
	for _, item := range items {
		out += "- " + item + "\n"
	}
	return out
}
