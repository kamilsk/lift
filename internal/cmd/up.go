package cmd

import "github.com/spf13/cobra"

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dump execution instructions based on configuration file for eval",
	Long:  "Dump execution instructions based on configuration file for eval.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
