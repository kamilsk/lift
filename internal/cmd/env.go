package cmd

import (
	"fmt"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/platform/pkg/unsafe"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Dump environment variables from configuration file",
	Long:  "Dump environment variables from configuration file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := config.FromFile(cmd.Flag("file").Value.String())
		if err != nil {
			return err
		}
		for env, value := range cnf.Environment {
			unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "%s=%s\n", env, value))
		}
		return nil
	},
}
