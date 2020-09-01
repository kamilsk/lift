package strings

// FirstNotEmpty returns a first non-empty string.
func FirstNotEmpty(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

// NotEmpty filters empty strings in-place.
func NotEmpty(strings []string) []string {
	filtered := strings[:0]
	for _, str := range strings {
		if str != "" {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

// Unique filters non-unique strings in-place.
func Unique(strings []string) []string {
	registry := map[string]struct{}{}
	filtered := strings[:0]
	for _, str := range strings {
		if _, present := registry[str]; present {
			continue
		}
		registry[str] = struct{}{}
		filtered = append(filtered, str)
	}
	return filtered
}
