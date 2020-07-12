package cmd

import "github.com/spf13/cobra"

func NewCallCommand() *cobra.Command {
	command := cobra.Command{
		Use:     "call",
		Short:   "todo",
		Long:    "TODO.",
		Example: "lift call ...",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return &command
}
