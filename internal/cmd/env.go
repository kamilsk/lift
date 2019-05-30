package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Dump environment variables from configuration file",
	Long:  "Dump environment variables from configuration file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := config.FromScope(scope(cmd))
		if err != nil {
			return err
		}
		vars := make([]string, 0, len(cnf.Environment))
		for variable, value := range forward.TransformEnvironment(cnf) {
			vars = append(vars, fmt.Sprintf("%s=%s", variable, value))
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout(), strings.Join(vars, "\n"))
		return err
	},
}
