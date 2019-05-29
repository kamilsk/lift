package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/shell"
)

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Execute another command with injecting environment variables into it",
	Long:  "Execute another command with injecting environment variables into it.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		cnf, err := config.FromFile(wd, cmd.Flag("file").Value.String())
		if err != nil {
			return err
		}
		vars := make([]string, 0, len(cnf.Environment))
		for variable, value := range cnf.Environment {
			vars = append(vars, fmt.Sprintf("%s=%s", variable, value))
		}
		var command = shell.Command(args[0])
		return shell.New(os.Getenv("SHELL")).Exec(command, args[1:], vars, cmd.OutOrStdout(), cmd.OutOrStderr())
	},
}
