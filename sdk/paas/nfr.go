package paas

import "sort"

// A NFR describes non-functional requirements of a service.
type NFR struct {
	Defaults Defaults  `toml:"default,omitempty"`
	Handlers []Handler `toml:"handlers"`
	Quota    []Token   `toml:"quota"`
}

// Merge combines two non-functional requirements of a service.
func (dst *NFR) Merge(nfr ...NFR) {
	if dst == nil || len(nfr) == 0 {
		return
	}

	for _, src := range nfr {

		// TODO:extend
		if len(src.Quota) > 0 {
			dst.Quota = append(dst.Quota, src.Quota...)
			sort.Slice(dst.Quota, func(i, j int) bool { return dst.Quota[i].ID < dst.Quota[j].ID })
		}

	}
}

// A Defaults contains default non-functional requirements of a service.
type Defaults struct {
	Handlers Handler `toml:"handlers"`
}

// A Handler contains configuration for request handler.
type Handler struct {
	Name        string `toml:"name"`
	Type        string `toml:"type"`
	Latency     `toml:",omitempty,squash"`
	Reliability `toml:",omitempty,squash"`
	Throughput  `toml:",omitempty,squash"`
}

// A Latency contains percentile values of response time.
type Latency struct {
	P75  string `toml:"latency_p75,omitempty"`
	P95  string `toml:"latency_p95,omitempty"`
	P98  string `toml:"latency_p98,omitempty"`
	P99  string `toml:"latency_p99,omitempty"`
	P999 string `toml:"latency_p999,omitempty"`
}

// A Reliability contains a value of error responses.
type Reliability struct {
	ErrorCodes   []uint `toml:"error_codes,omitempty"`
	ErrorPercent string `toml:"errors_percent,omitempty"`
}

// A Throughput contains a value of maximum throughput.
type Throughput struct {
	RPM string `toml:"max_rpm,omitempty"`
	RPS string `toml:"max_rps,omitempty"`
}

// A Token describes a consumer quota.
type Token struct {
	ID          string  `toml:"id" header:"x-source"`
	Consumer    string  `toml:"consumer"`
	Description string  `toml:"description"`
	Engine      string  `toml:"engine"`
	RequestedBy string  `toml:"requested_by"`
	Handlers    []Quota `toml:"handlers"`
}

// A Quota describes consumer limits of operation.
type Quota struct {
	Name        string `toml:"name"`
	Scope       string `toml:"scope" header:"x-scope"`
	Description string `toml:"description"`
	Latency     `toml:",omitempty,squash"`
	Throughput  `toml:",omitempty,squash"`
	Metadata    map[string]interface{} `toml:"metadata"`
}
