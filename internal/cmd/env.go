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
		unsafe.DoSilent(fmt.Fprintln(cmd.OutOrStdout(), cnf))
		return err
	},
}
