package paas

import (
	"fmt"
	"sort"
)

// A NFR describes non-functional requirements of a service.
type NFR struct {
	Defaults Defaults  `toml:"default,omitempty"`
	Handlers []Handler `toml:"handlers,omitempty"`
	Quota    []Token   `toml:"quota,omitempty"`
}

// Merge combines two non-functional requirements of a service.
func (dst *NFR) Merge(sources ...NFR) {
	if dst == nil || len(sources) == 0 {
		return
	}

	for _, src := range sources {
		dst.Defaults.Merge(src.Defaults)

		if len(src.Handlers) > 0 {
			dst.Handlers = append(dst.Handlers, src.Handlers...)
		}

		if len(src.Quota) > 0 {
			dst.Quota = append(dst.Quota, src.Quota...)
		}
	}

	if len(dst.Handlers) > 0 {
		sort.Slice(dst.Handlers, func(i, j int) bool {
			if dst.Handlers[i].Type == dst.Handlers[j].Type {
				return dst.Handlers[i].Name < dst.Handlers[j].Name
			}
			return dst.Handlers[i].Type < dst.Handlers[j].Type
		})
		registry := map[string]int{}
		for i := len(dst.Handlers); i > 0; i-- {
			registry[dst.Handlers[i-1].ID()] = i - 1
		}
		unique := dst.Handlers[:0]
		for i, handler := range dst.Handlers {
			origin := registry[handler.ID()]
			if i == origin {
				unique = append(unique, handler)
				continue
			}
			unique[origin].Merge(handler)
		}
		dst.Handlers = unique
	}

	if len(dst.Quota) > 0 {
		sort.Slice(dst.Quota, func(i, j int) bool { return dst.Quota[i].ID < dst.Quota[j].ID })
		registry := map[string]int{}
		for i := len(dst.Quota); i > 0; i-- {
			registry[dst.Quota[i-1].ID] = i - 1
		}
		unique := dst.Quota[:0]
		for i, token := range dst.Quota {
			origin := registry[token.ID]
			if i == origin {
				unique = append(unique, token)
				continue
			}
			unique[origin].Merge(token)
		}
		dst.Quota = unique
	}
}

// A Defaults contains default non-functional requirements of a service.
type Defaults struct {
	Handlers Handler `toml:"handlers,omitempty"`
}

// Merge combines two default non-functional requirements of a service.
func (dst *Defaults) Merge(src Defaults) {
	if dst == nil {
		return
	}

	dst.Handlers.Merge(src.Handlers)
}

// A Handler contains configuration for request handler.
type Handler struct {
	Name        string `toml:"name,omitempty"`
	Type        string `toml:"type,omitempty"`
	Latency     `toml:",omitempty,squash"`
	Reliability `toml:",omitempty,squash"`
	Throughput  `toml:",omitempty,squash"`
}

// ID returns a unique identifier for consumer limits of operation.
func (dst Handler) ID() string {
	return fmt.Sprintf("%s:%s", dst.Name, dst.Type)
}

// Merge combines two request handler configurations.
func (dst *Handler) Merge(src Handler) {
	if dst == nil || dst.ID() != src.ID() {
		return
	}

	dst.Latency.Merge(src.Latency)
	dst.Reliability.Merge(src.Reliability)
	dst.Throughput.Merge(src.Throughput)
}

// A Latency contains percentile values of response time.
type Latency struct {
	P75  string `toml:"latency_p75,omitempty"`
	P95  string `toml:"latency_p95,omitempty"`
	P98  string `toml:"latency_p98,omitempty"`
	P99  string `toml:"latency_p99,omitempty"`
	P999 string `toml:"latency_p999,omitempty"`
}

// Merge combines two percentile values of response time.
func (dst *Latency) Merge(src Latency) {
	if dst == nil {
		return
	}

	if src.P75 != "" {
		dst.P75 = src.P75
	}
	if src.P95 != "" {
		dst.P95 = src.P95
	}
	if src.P98 != "" {
		dst.P98 = src.P98
	}
	if src.P99 != "" {
		dst.P99 = src.P99
	}
	if src.P999 != "" {
		dst.P999 = src.P999
	}
}

// A Reliability contains a value of error responses.
type Reliability struct {
	ErrorCodes   []uint `toml:"error_codes,omitempty"`
	ErrorPercent string `toml:"errors_percent,omitempty"`
}

// Merge combines two values of error responses.
func (dst *Reliability) Merge(src Reliability) {
	if dst == nil {
		return
	}

	if len(src.ErrorCodes) > 0 {
		dst.ErrorCodes = append(dst.ErrorCodes, src.ErrorCodes...)
		sort.Slice(dst.ErrorCodes, func(i, j int) bool { return dst.ErrorCodes[i] < dst.ErrorCodes[j] })
		current, unique := dst.ErrorCodes[0], dst.ErrorCodes[:1]
		for _, code := range dst.ErrorCodes[1:] {
			if current != code {
				current = code
				unique = append(unique, current)
			}
		}
		dst.ErrorCodes = unique
	}

	if src.ErrorPercent != "" {
		dst.ErrorPercent = src.ErrorPercent
	}
}

// A Throughput contains a value of maximum throughput.
type Throughput struct {
	RPM string `toml:"max_rpm,omitempty"`
	RPS string `toml:"max_rps,omitempty"`
}

// Merge combines two values of maximum throughput.
func (dst *Throughput) Merge(src Throughput) {
	if dst == nil {
		return
	}

	if src.RPM != "" {
		dst.RPM = src.RPM
	}
	if src.RPS != "" {
		dst.RPS = src.RPS
	}
}

// A Token describes a consumer quota.
type Token struct {
	ID          string  `toml:"id" header:"x-source"`
	Consumer    string  `toml:"consumer,omitempty"`
	Description string  `toml:"description,omitempty"`
	Engine      string  `toml:"engine,omitempty"`
	RequestedBy string  `toml:"requested_by,omitempty"`
	Handlers    []Quota `toml:"handlers,omitempty"`
}

// Merge combines two consumer quota.
func (dst *Token) Merge(src Token) {
	if dst == nil || dst.ID != src.ID {
		return
	}

	if src.Consumer != "" {
		dst.Consumer = src.Consumer
	}
	if src.Description != "" {
		dst.Description = src.Description
	}
	if src.Engine != "" {
		dst.Engine = src.Engine
	}
	if src.RequestedBy != "" {
		dst.RequestedBy = src.RequestedBy
	}

	if len(src.Handlers) > 0 {
		copied := append(dst.Handlers, src.Handlers...)
		registry := map[string]int{}
		for i := len(copied); i > 0; i-- {
			registry[copied[i-1].ID()] = i - 1
		}
		unique := copied[:0]
		for i, quota := range copied {
			origin := registry[quota.ID()]
			if i == origin {
				unique = append(unique, quota)
				continue
			}
			unique[origin].Merge(quota)
		}
		dst.Handlers = unique
	}
}

// A Quota describes consumer limits of operation.
type Quota struct {
	Name        string `toml:"name"`
	Type        string `toml:"type,omitempty"`
	Scope       string `toml:"scope,omitempty" header:"x-scope"`
	Description string `toml:"description,omitempty"`
	Latency     `toml:",omitempty,squash"`
	Throughput  `toml:",omitempty,squash"`
	Metadata    map[string]interface{} `toml:"metadata,omitempty"`
}

// ID returns a unique identifier of consumer limits of operation.
func (dst Quota) ID() string {
	return fmt.Sprintf("%s:%s:%s", dst.Name, dst.Type, dst.Scope)
}

// Merge combines two consumer limits of operation.
func (dst *Quota) Merge(src Quota) {
	if dst == nil || dst.ID() != src.ID() {
		return
	}

	if src.Description != "" {
		dst.Description = src.Description
	}
	dst.Latency.Merge(src.Latency)
	dst.Throughput.Merge(src.Throughput)

	if dst.Metadata == nil && len(src.Metadata) > 0 {
		dst.Metadata = make(map[string]interface{})
	}
	for key, value := range src.Metadata {
		dst.Metadata[key] = value
	}
}
