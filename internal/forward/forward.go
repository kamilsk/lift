package forward

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/kamilsk/lift/internal/config"
	"github.com/kamilsk/lift/internal/shell"
)

var (
	port = regexp.MustCompile(`^\d+$`)
	host = regexp.MustCompile(`^(?:\w+.?)+:(\d+)$`)
)

// Command returns a command to run the forward tool for a service.
func Command(cnf config.Service, detach bool) (shell.Command, error) {
	var command shell.Command
	args := make([]string, 0, 8)
	for _, dep := range cnf.Dependencies {
		if len(dep.Forward) == 0 {
			continue
		}
		args = append(args, PodName(cnf.Name, dep.Name, true))
		for _, env := range dep.Forward {
			port, err := ExtractPort(cnf.Environment[env])
			if err != nil {
				return command, err
			}
			args = append(args, strconv.Itoa(int(port)))
		}
	}
	if len(args) > 0 {
		if detach {
			return shell.Command(fmt.Sprintf("forward -- %s &", strings.Join(args, " "))), nil
		}
		return shell.Command(fmt.Sprintf("forward -- %s", strings.Join(args, " "))), nil
	}
	return command, nil
}

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
		return uint16(p), nil
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
