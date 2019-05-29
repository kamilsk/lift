package main

import (
	"fmt"
	"os"

	"github.com/kamilsk/platform/cmd/cobra"

	"github.com/kamilsk/lift/internal/cmd"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() {
	root := cmd.New(nil)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(commit, date, version))
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
