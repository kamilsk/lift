package model

import "sort"

type Cron struct {
	Name     string `toml:"name,omitempty"`
	Enabled  *bool  `toml:"enabled,omitempty"`
	Schedule string `toml:"schedule,omitempty"`
	Command  string `toml:"command,omitempty"`
}

type Crons []Cron

func (crons Crons) Len() int           { return len(crons) }
func (crons Crons) Less(i, j int) bool { return crons[i].Name < crons[j].Name }
func (crons Crons) Swap(i, j int)      { crons[i], crons[j] = crons[j], crons[i] }

func (crons *Crons) Merge(src Crons) {
	if crons == nil || len(src) == 0 {
		return
	}

	copied := *crons
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*crons = copied[:shift+1]
}
