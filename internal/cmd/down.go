package cmd

import "github.com/spf13/cobra"

func NewDownCommand() *cobra.Command {
	command := cobra.Command{
		Use:     "down",
		Short:   "todo",
		Long:    "TODO.",
		Example: "lift down ...",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return &command
}
