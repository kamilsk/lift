package shell

import (
	"fmt"
	"io"
	"path"
	"reflect"
	"strings"
	"unsafe"
)

// Shell defines generic behavior of shells.
type Shell interface {
	// Assign returns a command to set the variable value.
	Assign(variable, value string) Command
	// Print dumps command list into the writer.
	Print(io.Writer, ...Command) error
}

// Command represents a valid shell command.
type Command string

// Shell returns a specific shell implementation.
func New(bin string) Shell {
	switch path.Base(bin) {
	case "sh":
		return sh{}
	case "ksh", "bash", "zsh":
		return bash{}
	case "csh":
		return csh{}
	case "cmd.exe":
		return win{}
	}
	panic(fmt.Errorf("cannot define related shell based on the passed bin %q", bin))
}

type sh struct{}

// Assign returns a command to set the variable value.
func (sh) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("%s=%q; export %[1]s", variable, value))
}

func (sh) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type bash struct{}

// Assign returns a command to set the variable value.
func (bash) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("export %s=%q", variable, value))
}

func (bash) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type csh struct{}

// Assign returns a command to set the variable value.
func (csh) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("setenv %s %q", variable, value))
}

func (csh) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type win struct{}

// Assign returns a command to set the variable value.
func (win) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("set %s=%q", variable, value))
}

func (win) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

func dump(output io.Writer, commands ...Command) error {
	head := *(*reflect.SliceHeader)(unsafe.Pointer(&commands))
	data := *(*[]string)(unsafe.Pointer(&head))
	_, err := fmt.Fprintln(output, strings.Join(data, ";\n"))
	return err
}
