// Package domain holds the core CDAD concepts the CLI operates on,
// independent of any command-line, filesystem, or ADE-specific concern.
package domain

// CDADProject is a project directory that CDAD structure is created in,
// or inspected from.
type CDADProject struct {
	RootPath string
}
