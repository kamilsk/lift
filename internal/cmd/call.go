package cmd

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.octolab.org/safe"

	"github.com/kamilsk/lift/internal/model"
)

func NewCallCommand() *cobra.Command {
	fs := afero.NewOsFs()

	command := cobra.Command{
		Use:     "call",
		Short:   "build app.toml from components",
		Long:    "Build app.toml from components.",
		Example: "lift call components/*.toml components/*/*.toml app.toml",
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// https://github.com/golang/go/issues/11862
			var matches []string
			for _, from := range args[:len(args)-1] {
				current, err := afero.Glob(fs, from)
				if err != nil {
					return errors.Wrapf(err, "find components for %q", from)
				}
				matches = append(matches, current...)
			}

			file, err := os.Create(args[len(args)-1])
			if err != nil {
				return errors.Wrap(err, "create output file")
			}
			defer safe.Close(file, func(err error) { panic(err) })

			var app model.Application
			for _, component := range matches {
				var src model.Application

				tree, err := toml.LoadFile(component)
				if err != nil {
					return errors.Wrapf(err, "load component %q", component)
				}

				decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
					Result:  &src,
					TagName: "toml",
				})
				if err != nil {
					return errors.Wrap(err, "create decoder")
				}

				if err := decoder.Decode(tree.ToMap()); err != nil {
					return errors.Wrapf(err, "decode component %q", component)
				}
				app.Merge(src)
			}

			return errors.Wrap(toml.NewEncoder(file).Encode(app), "write output file")
		},
	}
	return &command
}
