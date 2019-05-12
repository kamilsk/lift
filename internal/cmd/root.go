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
	cmd.AddCommand(envCmd, upCmd)
	cmd.PersistentFlags().String("file", "app.toml", "service configuration file")
	cmd.SetOutput(output)
	return &cmd
}
