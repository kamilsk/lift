package cmd

import "github.com/spf13/cobra"

// New returns new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "lift",
		Short: "up service locally",
		Long:  "Up service locally.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	command.AddCommand(upCmd, downCmd, envCmd, forwardCmd, callCmd)
	flags := command.PersistentFlags()
	flags.StringP("file", "f", "app.toml", "service configuration file")
	flags.StringArrayP("map", "m", nil, "port mapping (e.g. -m REMOTE:LOCAL)")
	return &command
}
