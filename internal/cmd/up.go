package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/cnf"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dump instruction for eval to up environment locally",
	Long:  "Dump instruction for eval to up environment locally.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cnf.FromScope(scope(cmd))
		if err != nil {
			return err
		}
		sh := shell.New(os.Getenv("SHELL"))
		commands := make([]shell.Command, 0, 8)
		for variable, value := range forward.TransformEnvironment(config) {
			commands = append(commands, sh.Assign(variable, value))
		}
		return sh.Print(cmd.OutOrStdout(), commands...)
	},
}
