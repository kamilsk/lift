package cobra

import (
	"fmt"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
	"go.octolab.org/toolkit/config"
)

// NewVersionCommand returns a command that helps to build version info.
//
//  $ cli version
//  cli:
//    version     : 1.0.0
//    build date  : 2019-07-17T12:44:00Z
//    git hash    : 4f8c7f4
//    go version  : go1.12.7
//    go compiler : gc
//    platform    : darwin/amd64
//    features    : featureA=true, featureB=false
//
func NewVersionCommand(release, date, hash string, features ...config.Feature) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show application version",
		Long:  "Show application version.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return version.Execute(cmd.OutOrStdout(), struct {
				Name       string
				Version    string
				BuildDate  string
				GitHash    string
				GoVersion  string
				GoCompiler string
				Platform   string
				Features   fmt.Stringer
			}{
				Name:       root(cmd).Name(),
				Version:    release,
				BuildDate:  date,
				GitHash:    hash,
				GoVersion:  runtime.Version(),
				GoCompiler: runtime.Compiler,
				Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
				Features:   config.Features(features),
			})
		},
	}
}

var version = template.Must(template.New("version").Parse(`{{.Name}}:
  version     : {{.Version}}
  build date  : {{.BuildDate}}
  git hash    : {{.GitHash}}
  go version  : {{.GoVersion}}
  go compiler : {{.GoCompiler}}
  platform    : {{.Platform}}
  features    : {{.Features}}
`))
