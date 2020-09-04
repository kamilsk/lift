package model

// A SFTP contains configuration for a sftp server.
type SFTP struct {
	Enabled *bool  `toml:"enabled"`
	Size    string `toml:"size"`
}

// Merge combines two sftp server configurations.
func (dst *SFTP) Merge(src *SFTP) {
	if dst == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Size != "" {
		dst.Size = src.Size
	}
}
