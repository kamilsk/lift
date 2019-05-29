package cmd

import (
	"fmt"
	"os"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/shell"
	"github.com/spf13/cobra"
)

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call another command with injecting environment variables",
	Long:  "Call another command with injecting environment variables.",
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
