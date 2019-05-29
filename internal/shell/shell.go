package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"
	"unsafe"
)

// Shell defines generic behavior of shells.
type Shell interface {
	// Assign returns a command to set the variable value.
	Assign(variable, value string) Command
	// Exec executes the command with the specified arguments and environment variables.
	Exec(command Command, args, vars []string, stdout, stderr io.Writer) error
	// Print dumps command list into the writer.
	Print(io.Writer, ...Command) error
}

// Command represents a valid shell command.
type Command string

// Shell returns a specific shell implementation.
func New(bin string) Shell {
	switch path.Base(bin) {
	case "sh":
		return sh{bin}
	case "ksh", "bash", "zsh":
		return bash{bin}
	case "csh":
		return csh{bin}
	case "cmd.exe":
		return win{bin}
	}
	panic(fmt.Errorf("cannot define related shell based on the passed bin %q", bin))
}

type sh struct{ bin string }

// Assign returns a command to set the variable value.
func (sh) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("%s=%s; export %[1]s", variable, value))
}

// Exec executes the command with the specified arguments and environment variables.
func (sh) Exec(command Command, args, vars []string, stdout, stderr io.Writer) error {
	return nil
}

func (sh) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type bash struct{ bin string }

// Assign returns a command to set the variable value.
func (bash) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("export %s=%s", variable, value))
}

// Exec executes the command with the specified arguments and environment variables.
func (bash bash) Exec(command Command, args, vars []string, stdout, stderr io.Writer) error {
	args = append([]string{"-c", string(command)}, args...)
	args = append(args[:1], strings.Join(args[1:], " "))
	return execute(bash.bin, args, vars, stdout, stderr)
}

func (bash) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type csh struct{ bin string }

// Assign returns a command to set the variable value.
func (csh) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("setenv %s %s", variable, value))
}

// Exec executes the command with the specified arguments and environment variables.
func (csh) Exec(command Command, args, vars []string, stdout, stderr io.Writer) error {
	return nil
}

func (csh) Print(output io.Writer, commands ...Command) error {
	return dump(output, commands...)
}

type win struct{ bin string }

// Assign returns a command to set the variable value.
func (win) Assign(variable, value string) Command {
	return Command(fmt.Sprintf("set %s=%s", variable, value))
}

// Exec executes the command with the specified arguments and environment variables.
func (win) Exec(command Command, args, vars []string, stdout, stderr io.Writer) error {
	return nil
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

func execute(shell string, args, vars []string, stdout, stderr io.Writer) error {
	cmd := exec.Command(shell, args...)
	cmd.Env = append(os.Environ(), vars...)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	return cmd.Run()
}
