package cobra

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// After inserts a new function into the pointer, which calls the self function before and the last after.
func After(pointer *func(*cobra.Command, []string), last func(*cobra.Command, []string)) {
	first := *pointer
	if first == nil {
		first = func(*cobra.Command, []string) {}
	}
	*pointer = func(command *cobra.Command, args []string) {
		first(command, args)
		last(command, args)
	}
}

// AfterE inserts a new function into the pointer, which calls the self function before and the last after.
func AfterE(pointer *func(*cobra.Command, []string) error, last func(*cobra.Command, []string) error) {
	first := *pointer
	if first == nil {
		first = func(*cobra.Command, []string) error { return nil }
	}
	*pointer = func(command *cobra.Command, args []string) error {
		if err := first(command, args); err != nil {
			return err
		}
		return last(command, args)
	}
}

// Apply applies options to the Command.
func Apply(command *cobra.Command, container *viper.Viper, options ...Option) *cobra.Command {
	for _, configure := range options {
		configure(command, container)
	}
	return command
}

// Before inserts a new function into the pointer, which calls the first function before and the self after.
func Before(pointer *func(*cobra.Command, []string), first func(*cobra.Command, []string)) {
	last := *pointer
	if last == nil {
		last = func(*cobra.Command, []string) {}
	}
	*pointer = func(command *cobra.Command, args []string) {
		first(command, args)
		last(command, args)
	}
}

// BeforeE inserts a new function into the pointer, which calls the first function before and the self after.
func BeforeE(pointer *func(*cobra.Command, []string) error, first func(*cobra.Command, []string) error) {
	last := *pointer
	if last == nil {
		last = func(*cobra.Command, []string) error { return nil }
	}
	*pointer = func(command *cobra.Command, args []string) error {
		if err := first(command, args); err != nil {
			return err
		}
		return last(command, args)
	}
}

// An Option is a Command configuration function.
type Option func(*cobra.Command, *viper.Viper)
