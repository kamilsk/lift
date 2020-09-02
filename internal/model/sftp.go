package model

// A SFTP contains configuration for a sftp server.
type SFTP struct {
	Enabled *bool  `toml:"enabled"`
	Size    string `toml:"size"`
}

// Merge combines two sftp server configurations.
func (sftp *SFTP) Merge(src *SFTP) {
	if sftp == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		sftp.Enabled = src.Enabled
	}
	if src.Size != "" {
		sftp.Size = src.Size
	}
}
