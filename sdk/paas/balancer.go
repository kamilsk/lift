package paas

// A Balancer contains configuration for a balancer.
type Balancer struct {
	CookieAffinity string `toml:"cookie_affinity"`
}

// Merge combines two balancer configurations.
func (dst *Balancer) Merge(src *Balancer) {
	if dst == nil || src == nil {
		return
	}

	if src.CookieAffinity != "" {
		dst.CookieAffinity = src.CookieAffinity
	}
}
