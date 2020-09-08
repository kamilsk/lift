package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "lift",
		Short: "up your service locally",
		Long:  "Up your service locally.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	command.AddCommand(
		NewCallCommand(),
	)
	return &command
}
