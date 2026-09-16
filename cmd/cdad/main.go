package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/CDAD-Community/cdad-cli/internal/bootstrap"
	"github.com/CDAD-Community/cdad-cli/internal/domain"
	"github.com/CDAD-Community/cdad-cli/internal/validator"
)

// version is overridden at build time via -ldflags "-X main.version=vX.Y.Z".
var version = "dev"

func main() {
	root := &cobra.Command{
		Use:   "cdad",
		Short: "CDAD CLI — govern AI-assisted software development workflows",
	}

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the cdad CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), version)
			return nil
		},
	})

	root.AddCommand(newInitCommand())
	root.AddCommand(newValidateCommand())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func newInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init [path]",
		Short: "Create the CDAD structure in a project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) == 1 {
				path = args[0]
			}
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			report, err := bootstrap.Init(domain.CDADProject{RootPath: abs}, version)
			if err != nil {
				return err
			}

			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "cdad init -- %s\n\n", abs)
			printSection(out, "Created", report.Created)
			printSection(out, "Conflicts (already existed, left untouched)", report.Conflicts)
			printSection(out, "Pending (bootstrap artifact acquisition not yet implemented)", report.Pending)
			return nil
		},
	}
}

func newValidateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "validate [path]",
		Short: "Run deterministic governance checks (no LLM required)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) == 1 {
				root = args[0]
			}
			backlogPath := filepath.Join(root, "cdad", "backlog.md")
			out := cmd.OutOrStdout()

			report, err := validator.ValidateBacklog(backlogPath)
			if os.IsNotExist(err) {
				fmt.Fprintf(out, "cdad validate: %s not found.\n", backlogPath)
				os.Exit(2)
			}
			if err != nil {
				return err
			}

			printBacklogReport(out, backlogPath, report)

			if report.Failed() {
				os.Exit(1)
			}
			return nil
		},
	}
}

func printBacklogReport(out io.Writer, path string, report *validator.BacklogReport) {
	fmt.Fprintf(out, "cdad validate -- %s\n\n", path)
	printSection(out, "Duplicate Epic IDs", report.DuplicateEpicIDs)
	printSection(out, "Duplicate Story IDs", report.DuplicateStoryIDs)
	printSection(out, "Epic statuses outside the agreed vocabulary", report.BadEpicStatuses)
	printSection(out, "Story statuses outside the agreed vocabulary", report.BadStoryStatuses)

	if len(report.EmptyEpics) > 0 {
		fmt.Fprintln(out, "Warnings (Epics with no Stories yet):")
		for _, id := range report.EmptyEpics {
			fmt.Fprintf(out, "  - %s\n", id)
		}
		fmt.Fprintln(out)
	}

	if report.Failed() {
		fmt.Fprintln(out, "cdad validate: FAILED")
	} else {
		fmt.Fprintln(out, "cdad validate: OK")
	}
}

func printSection(out io.Writer, title string, items []string) {
	fmt.Fprintf(out, "%s:\n", title)
	if len(items) == 0 {
		fmt.Fprintln(out, "  none")
		return
	}
	for _, item := range items {
		fmt.Fprintf(out, "  - %s\n", item)
	}
	fmt.Fprintln(out)
}
