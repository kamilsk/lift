package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/config"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Dump environment variables from configuration file",
	Long:  "Dump environment variables from configuration file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := scope(cmd)
		if err != nil {
			return err
		}
		cnf, err := config.FromScope(ctx)
		if err != nil {
			return err
		}
		vars := make([]string, 0, len(cnf.Environment))
		for variable, value := range cnf.Environment {
			vars = append(vars, fmt.Sprintf("%s=%s", variable, value))
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout(), strings.Join(vars, "\n"))
		return err
	},
}
