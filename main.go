package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kamilsk/lift/internal/cmd"
	"github.com/kamilsk/platform/cmd/cobra"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		cancel()
		time.Sleep(50 * time.Millisecond)
		signal.Stop(c)
		fmt.Println()
		os.Exit(0)
	}()

	root := cmd.New(nil)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(commit, date, version))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
