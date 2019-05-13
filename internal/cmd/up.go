package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/forward"
	"github.com/kamilsk/platform/pkg/unsafe"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dump execution instructions based on configuration file for eval",
	Long:  "Dump execution instructions based on configuration file for eval.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cnf, err := config.FromFile(cmd.Flag("file").Value.String())
		if err != nil {
			return err
		}
		for env, value := range cnf.Environment {
			unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "export %s=%q;\n", env, value)) // TODO:batch flush together
		}
		{
			args := make([]string, 0, 8)
			for _, dep := range cnf.Dependencies {
				if len(dep.Forward) == 0 {
					continue
				}
				args = append(args, forward.PodName(cnf.Name, dep.Name, true))
				for _, env := range dep.Forward {
					port, err := forward.ExtractPort(cnf.Environment[env])
					if err != nil {
						return err
					}
					args = append(args, strconv.Itoa(int(port)))
				}
			}
			if len(args) > 0 {
				// TODO:upstream use demonized version with signal support
				unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "forward -- %s &;\n", strings.Join(args, " ")))
			}
		}
		if len(args) == 0 {
			args = []string{"cmd/service/main.go"} // TODO:check use os.Stat to filter valid entry
		}
		unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "go run %s;\n", strings.Join(args, " ")))
		unsafe.DoSilent(fmt.Fprintf(cmd.OutOrStdout(), "ps | grep '[f]orward --' | awk '{print $1}' | xargs kill -SIGKILL;\n"))
		return nil
	},
}
