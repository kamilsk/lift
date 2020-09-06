package paas

import (
	"strings"

	xstrings "go.octolab.org/strings"
)

// A RabbitMQ contains configuration for a database.
type RabbitMQ struct {
	Enabled *bool  `toml:"enabled"`
	Version string `toml:"version"`
	Size    string `toml:"size"`
	Vhosts  string `toml:"vhosts,omitempty"`
}

// Merge combines two database configurations.
func (dst *RabbitMQ) Merge(src *RabbitMQ) {
	if dst == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Version != "" {
		dst.Version = src.Version
	}
	if src.Size != "" {
		dst.Size = src.Size
	}

	if src.Vhosts != "" {
		dst.Vhosts = strings.Join(
			xstrings.Unique(
				xstrings.NotEmpty(
					strings.Split(dst.Vhosts+","+src.Vhosts, ","),
				),
			),
			",",
		)
	}
}
