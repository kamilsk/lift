package forward

import "strings"

// PodName builds pod name.
func PodName(service, entity string, isLocal bool) string {
	parts := make([]string, 0, 4)
	parts = append(parts, service)
	if isLocal {
		parts = append(parts, "local")
	}
	parts = append(parts, entity, "")
	return strings.ToLower(strings.Join(parts, "-"))
}
