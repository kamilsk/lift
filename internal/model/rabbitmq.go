package model

type RabbitMQ struct {
	Enabled *bool  `toml:"enabled,omitempty"`
	Version string `toml:"version,omitempty"`
	Size    string `toml:"size,omitempty"`
	Vhosts  string `toml:"vhosts,omitempty"`
}

func (rabbitmq *RabbitMQ) Merge(src *RabbitMQ) {
	if rabbitmq == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		rabbitmq.Enabled = src.Enabled
	}
	if src.Version != "" {
		rabbitmq.Version = src.Version
	}
	if src.Size != "" {
		rabbitmq.Size = src.Size
	}
	if src.Vhosts != "" {
		rabbitmq.Vhosts = src.Vhosts
	}
}
