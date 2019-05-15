package forward

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
