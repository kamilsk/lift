package model

// A Logger contains configuration for a logger.
type Logger struct {
	Level string `toml:"level"`
}

// Merge combines two logger configurations.
func (dst *Logger) Merge(src *Logger) {
	if dst == nil || src == nil {
		return
	}

	if src.Level != "" {
		dst.Level = src.Level
	}
}
