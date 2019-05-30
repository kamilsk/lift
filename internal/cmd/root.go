package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

// New returns new root command.
func New(output io.Writer) *cobra.Command {
	cmd := cobra.Command{
		Use:   "lift",
		Short: "Up service locally",
		Long:  "Up service locally.",
	}
	cmd.AddCommand(upCmd, downCmd, envCmd, forwardCmd, callCmd)
	cmd.SetOutput(output)
	var (
		file string
	)
	cmd.PersistentFlags().StringVarP(&file, "file", "f", "app.toml", "service configuration file")
	return &cmd
}
