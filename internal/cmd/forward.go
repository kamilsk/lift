package cmd

import (
	"os"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
	"github.com/spf13/cobra"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Dump instruction for port forwarding",
	Long:  "Dump instruction for port forwarding.",
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		cnf, err := config.FromFile(wd, cmd.Flag("file").Value.String())
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
