package forward

import (
	"net/url"
	"regexp"
	"strconv"
)

var (
	digit = regexp.MustCompile(`^\d+$`)
	host  = regexp.MustCompile(`^(?:\w+.?)+:(\d+)$`)
)

// ExtractPort tries to extract a port from a connection definition.
func ExtractPort(definition string) (uint16, error) {
	switch {
	case digit.MatchString(definition):
		port, err := strconv.ParseUint(definition, 10, 16)
		return uint16(port), err
	case host.MatchString(definition):
		result := host.FindStringSubmatch(definition)
		port, err := strconv.ParseUint(result[len(result)-1], 10, 16)
		return uint16(port), err
	default:
		u, err := url.Parse(definition)
		if err != nil {
			return 0, err
		}
		port, err := strconv.ParseUint(u.Port(), 10, 16)
		return uint16(port), nil
	}
}
