package paas

// A Cron contains configuration for a cron job.
type Cron struct {
	Name      string     `toml:"name"`
	Enabled   *bool      `toml:"enabled"`
	Schedule  string     `toml:"schedule"`
	Command   string     `toml:"command"`
	Resources *Resources `toml:"resources,omitempty"`
}

// Merge combines two cron job configurations.
func (dst *Cron) Merge(src Cron) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Schedule != "" {
		dst.Schedule = src.Schedule
	}
	if src.Command != "" {
		dst.Command = src.Command
	}

	if src.Resources != nil && dst.Resources == nil {
		dst.Resources = new(Resources)
	}
	dst.Resources.Merge(src.Resources)
}

// Crons is a list of Cron.
type Crons []Cron

// Len, Less, Swap implements the sort.Interface.
func (dst Crons) Len() int           { return len(dst) }
func (dst Crons) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Crons) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of cron job configurations.
func (dst *Crons) Merge(src Crons) {
	if dst == nil || len(src) == 0 {
		return
	}

	copied := *dst
	copied = append(copied, src...)

	registry := map[string]int{}
	for i := len(copied); i > 0; i-- {
		registry[copied[i-1].Name] = i - 1
	}
	unique := copied[:0]
	for i, cron := range copied {
		origin := registry[cron.Name]
		if i == origin {
			unique = append(unique, cron)
			continue
		}
		unique[origin].Merge(cron)
	}

	*dst = unique
}
