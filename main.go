package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	_ "github.com/pkg/errors"
	_ "github.com/spf13/afero"
	"go.octolab.org/errors"
	"go.octolab.org/safe"
	"go.octolab.org/toolkit/cli/cobra"
	"go.octolab.org/unsafe"

	"github.com/kamilsk/lift/internal/cmd"
	"github.com/kamilsk/lift/internal/cnf"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
	exit    = os.Exit
	stderr  = io.Writer(os.Stderr)
	stdout  = io.Writer(os.Stdout)
)

func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	root := cmd.New()
	root.SetErr(stderr)
	root.SetOut(stdout)
	root.AddCommand(
		cobra.NewCompletionCommand(),
		cobra.NewVersionCommand(version, date, commit, cnf.Features...),
	)
	safe.Do(func() error { return root.ExecuteContext(ctx) }, shutdown)
}

func shutdown(err error) {
	if recovered, is := errors.Unwrap(err).(errors.Recovered); is {
		unsafe.DoSilent(fmt.Fprintf(stderr, "recovered: %+v\n", recovered.Cause()))
		unsafe.DoSilent(fmt.Fprintln(stderr, "---"))
		unsafe.DoSilent(fmt.Fprintf(stderr, "%+v\n", err))
	}
	exit(1)
}
