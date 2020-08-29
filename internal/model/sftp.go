package model

type SFTP struct {
	Size    string `toml:"size,omitempty"`
	Enabled *bool  `toml:"enabled,omitempty"`
}

func (sftp *SFTP) Merge(src *SFTP) {
	if sftp == nil || src == nil {
		return
	}

	if src.Size != "" {
		sftp.Size = src.Size
	}
	if src.Enabled != nil {
		sftp.Enabled = src.Enabled
	}
}
