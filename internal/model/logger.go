package model

type Logger struct {
	Level string `toml:"level,omitempty"`
}

func (logger *Logger) Merge(src *Logger) {
	if logger == nil || src == nil {
		return
	}

	if src.Level != "" {
		logger.Level = src.Level
	}
}
