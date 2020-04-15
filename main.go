package main

import (
	"fmt"
	"os"

	"go.octolab.org/toolkit/cli/cobra"

	"github.com/kamilsk/lift/internal/cmd"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func main() {
	root := cmd.New(nil)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(version, date, commit))
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
