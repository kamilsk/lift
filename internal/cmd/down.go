package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Dump instruction for eval to down environment locally",
	Long:  "Dump instruction for eval to down environment locally.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := scope(cmd)
		if err != nil {
			return err
		}
		cnf, err := config.FromScope(ctx)
		if err != nil {
			return err
		}
		sh := shell.New(os.Getenv("SHELL"))
		commands := make([]shell.Command, 0, 2)
		command, err := forward.Command(cnf, true)
		if err != nil {
			return err
		}
		if command != "" {
			commands = append(commands, forward.Shutdown(cnf)...)
		}
		return sh.Print(cmd.OutOrStdout(), commands...)
	},
}
