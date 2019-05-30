package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Dump instruction for port forwarding",
	Long:  "Dump instruction for port forwarding.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := scope(cmd)
		if err != nil {
			return err
		}
		cnf, err := config.FromScope(ctx)
		if err != nil {
			return err
		}
		command, err := forward.Command(cnf, false)
		if err != nil {
			return err
		}
		return shell.New(os.Getenv("SHELL")).Print(cmd.OutOrStdout(), command)
	},
}
