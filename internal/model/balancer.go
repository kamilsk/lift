package model

type Balancer struct {
	CookieAffinity string `toml:"cookie_affinity,omitempty"`
}

func (balancer *Balancer) Merge(src *Balancer) {
	if balancer == nil || src == nil {
		return
	}

	if src.CookieAffinity != "" {
		balancer.CookieAffinity = src.CookieAffinity
	}
}
