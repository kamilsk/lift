package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/lift/internal/shell"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dump instruction for eval to run service locally",
	Long:  "Dump instruction for eval to run service locally.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := config.FromFile(cmd.Flag("file").Value.String())
		if err != nil {
			return err
		}
		sh := shell.New(os.Getenv("SHELL"))
		commands := make([]shell.Command, 0, 8)
		for variable, value := range cnf.Environment {
			commands = append(commands, sh.Assign(variable, value))
		}
		command, err := forward.Command(cnf, true)
		if err != nil {
			return err
		}
		if command != "" {
			commands = append(commands, command)
		}
		if len(args) == 0 {
			args = []string{"cmd/service/main.go"} // TODO:check use os.Stat to filter valid entry
		}
		commands = append(commands, shell.Command(fmt.Sprintf("go run %s", strings.Join(args, " "))))
		commands = append(commands, shell.Command("ps | grep '[f]orward --' | awk '{print $1}' | xargs kill -SIGKILL"))
		return sh.Print(cmd.OutOrStdout(), commands...)
	},
}
