package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/cnf"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
)

var callCmd = &cobra.Command{
	Use:     "call",
	Short:   "Execute another command with injecting environment variables into it",
	Long:    "Execute another command with injecting environment variables into it.",
	Example: "lift call -- echo $GOMODULE",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cnf.FromScope(scope(cmd))
		if err != nil {
			return err
		}
		vars := make([]string, 0, len(config.Environment))
		for variable, value := range forward.TransformEnvironment(config) {
			vars = append(vars, fmt.Sprintf("%s=%s", variable, value))
		}
		var command = shell.Command(args[0])
		return shell.New(os.Getenv("SHELL")).Exec(command, args[1:], vars, cmd.OutOrStdout(), cmd.OutOrStderr())
	},
}
