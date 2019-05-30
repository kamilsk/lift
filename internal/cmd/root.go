package cmd

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/kamilsk/lift/internal"
)

// New returns new root command.
func New(output io.Writer) *cobra.Command {
	cmd := cobra.Command{
		Use:   "lift",
		Short: "Up service locally",
		Long:  "Up service locally.",
	}
	cmd.AddCommand(upCmd, downCmd, envCmd, forwardCmd, callCmd)
	cmd.SetOutput(output)
	var (
		file    string
		mapping = make([]string, 0, 4)
	)
	cmd.PersistentFlags().StringVarP(&file, "file", "f", "app.toml", "service configuration file")
	cmd.PersistentFlags().StringArrayVarP(&mapping, "map", "m", nil, "port mapping (e.g. -m REMOTE:LOCAL)")
	return &cmd
}

func scope(cmd *cobra.Command) (internal.Scope, error) {
	var scope = internal.Scope{PortMapping: make(map[uint16]uint16)}

	wd, err := os.Getwd()
	if err != nil {
		return scope, err
	}
	mm, err := cmd.Flags().GetStringArray("map")
	if err != nil {
		return scope, err
	}
	for _, m := range mm {
		rl := strings.Split(m, ":")
		if len(rl) != 2 {
			return scope, errors.Errorf("unexpected port mapping, format REMOTE:LOCAL is expected, obtained %q", m)
		}
		remote, err := strconv.ParseUint(rl[0], 10, 16)
		if err != nil {
			return scope, err
		}
		local, err := strconv.ParseUint(rl[1], 10, 16)
		if err != nil {
			return scope, err
		}
		scope.PortMapping[uint16(remote)] = uint16(local)
	}
	scope.ConfigPath = cmd.Flag("file").Value.String()
	scope.WorkingDir = wd

	return scope, nil
}
