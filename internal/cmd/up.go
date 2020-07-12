package cmd

import "github.com/spf13/cobra"

func NewUpCommand() *cobra.Command {
	command := cobra.Command{
		Use:     "up",
		Short:   "todo",
		Long:    "TODO.",
		Example: "lift up ...",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return &command
}
