package model

// A Logger contains configuration for a logger.
type Logger struct {
	Level string `toml:"level"`
}

// Merge combines two logger configurations.
func (logger *Logger) Merge(src *Logger) {
	if logger == nil || src == nil {
		return
	}

	if src.Level != "" {
		logger.Level = src.Level
	}
}
