package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/cnf"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Dump instruction for eval to down environment locally",
	Long:  "Dump instruction for eval to down environment locally.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cnf.FromScope(scope(cmd))
		if err != nil {
			return err
		}
		sh := shell.New(os.Getenv("SHELL"))
		commands := make([]shell.Command, 0, 2)
		command, err := forward.Command(config, true)
		if err != nil {
			return err
		}
		if command != "" {
			commands = append(commands, forward.Shutdown(config)...)
		}
		return sh.Print(cmd.OutOrStdout(), commands...)
	},
}
