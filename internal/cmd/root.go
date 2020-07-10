package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "lift",
		Short: "up service locally",
		Long:  "Up service locally.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	flags := command.PersistentFlags()
	flags.StringP("file", "f", "app.toml", "service configuration file")
	flags.StringArrayP("map", "m", nil, "port mapping (e.g. -m REMOTE:LOCAL)")
	command.AddCommand(upCmd, downCmd, envCmd, callCmd)
	return &command
}
