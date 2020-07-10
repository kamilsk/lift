package forward

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/kamilsk/lift/internal/cnf"
	"github.com/kamilsk/lift/internal/shell"
)

var (
	port = regexp.MustCompile(`^\d+$`)
	host = regexp.MustCompile(`^(?:\w+.?)+:(\d+)$`)
)

// ExtractPort tries to extract a port number from a connection definition.
func ExtractPort(connection string) (uint16, error) {
	switch {
	case port.MatchString(connection):
		p, err := strconv.ParseUint(connection, 10, 16)
		return uint16(p), err
	case host.MatchString(connection):
		result := host.FindStringSubmatch(connection)
		p, err := strconv.ParseUint(result[len(result)-1], 10, 16)
		return uint16(p), err
	default:
		u, err := url.Parse(connection)
		if err != nil {
			return 0, err
		}
		p, err := strconv.ParseUint(u.Port(), 10, 16)
		return uint16(p), err
	}
}

// PodName builds pod name.
func PodName(service, entity string, isLocal bool) string {
	parts := append(make([]string, 0, 4), service)
	if isLocal {
		parts = append(parts, "local")
	}
	parts = append(parts, entity, "")
	return strings.ToLower(strings.Join(parts, "-"))
}

// ReplacePort replaces a port number in a connection definition.
func ReplacePort(connection string, from, to uint16) string {
	return strings.Replace(connection, strconv.FormatUint(uint64(from), 10), strconv.FormatUint(uint64(to), 10), 1)
}

// Shutdown returns commands to shutdown the forward tool.
func Shutdown(config cnf.Service) []shell.Command {
	return []shell.Command{
		"ps | grep '[f]orward --' | awk '{print $1}' | xargs kill -SIGKILL || true",
		shell.Command(
			fmt.Sprintf(
				"ps | grep '[f]orward %s' | awk '{print $1}' | xargs kill -SIGKILL || true",
				strings.TrimRight(PodName(config.Name, "", true), "-"),
			),
		),
	}
}

// TransformEnvironment applies a port mapping to a copy of environment variables.
func TransformEnvironment(config cnf.Service) cnf.Environment {
	if len(config.PortMapping) == 0 {
		return config.Environment
	}

	copied := make(cnf.Environment)
	for k, v := range config.Environment {
		copied[k] = v
	}

	for _, dep := range config.Dependencies {
		if len(dep.Forward) == 0 {
			continue
		}
		for _, env := range dep.Forward {
			remote, err := ExtractPort(copied[env])
			if err != nil {
				continue
			}
			if local, found := config.PortMapping[remote]; found {
				copied[env] = ReplacePort(copied[env], remote, local)
			}
		}
	}

	return copied
}
