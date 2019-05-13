package cmd

import (
	"fmt"
	"strings"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/platform/pkg/unsafe"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dump execution instructions based on configuration file for eval",
	Long:  "Dump execution instructions based on configuration file for eval.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := config.FromFile(cmd.Flag("file").Value.String())
		if err != nil {
			return err
		}
		for env, value := range cnf.Environment {
			unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "export %s=%q;\n", env, value))
		}
		if len(args) == 0 {
			args = []string{"main.go"}
		}
		unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "go run %s;\n", strings.Join(args, " ")))
		return nil
	},
}
