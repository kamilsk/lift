package cmd

import "github.com/spf13/cobra"

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Dump environment variables from configuration file",
	Long:  "Dump environment variables from configuration file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
