package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestNFR_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *NFR
		assert.NotPanics(t, func() { dst.Merge(NFR{Defaults: Defaults{Handlers: Handler{Name: "any"}}}) })
		assert.Nil(t, dst)
	})

	t.Run("no sources", func(t *testing.T) {
		var dst = new(NFR)
		assert.NotPanics(t, func() { dst.Merge() })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := NFR{
			Defaults: Defaults{Handlers: Handler{Reliability: Reliability{ErrorCodes: []uint{500}}}},
			Handlers: []Handler{
				{
					Name:       "handler",
					Type:       "brief",
					Latency:    Latency{P99: "100ms"},
					Throughput: Throughput{RPM: "100k"},
				},
			},
			Quota: []Token{
				{
					ID:          "token-a",
					Consumer:    "service-a",
					Engine:      "php",
					RequestedBy: "robot",
				},
			},
		}
		src := NFR{
			Defaults: Defaults{Handlers: Handler{Reliability: Reliability{ErrorCodes: []uint{499, 500, 503}}}},
			Handlers: []Handler{
				{
					Name:       "handler",
					Type:       "rest",
					Throughput: Throughput{RPM: "100k"},
				},
				{
					Name:       "handler",
					Type:       "brief",
					Latency:    Latency{P999: "200ms"},
					Throughput: Throughput{RPM: "500k"},
				},
			},
			Quota: []Token{
				{
					ID:          "token-a",
					Consumer:    "service-a",
					Description: "~",
					Engine:      "go",
					RequestedBy: "robot",
				},
				{
					ID:          "token-b",
					Consumer:    "service-a",
					Description: "~",
					Engine:      "go",
					RequestedBy: "robot",
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, NFR{
			Defaults: Defaults{Handlers: Handler{Reliability: Reliability{ErrorCodes: []uint{499, 500, 503}}}},
			Handlers: []Handler{
				{
					Name:       "handler",
					Type:       "brief",
					Latency:    Latency{P99: "100ms", P999: "200ms"},
					Throughput: Throughput{RPM: "500k"},
				},
				{
					Name:       "handler",
					Type:       "rest",
					Throughput: Throughput{RPM: "100k"},
				},
			},
			Quota: []Token{
				{
					ID:          "token-a",
					Consumer:    "service-a",
					Description: "~",
					Engine:      "go",
					RequestedBy: "robot",
				},
				{
					ID:          "token-b",
					Consumer:    "service-a",
					Description: "~",
					Engine:      "go",
					RequestedBy: "robot",
				},
			},
		}, dst)
	})
}

func TestDefaults_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Defaults
		assert.NotPanics(t, func() { dst.Merge(Defaults{Handlers: Handler{Name: "handler"}}) })
		assert.Nil(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Defaults{
			Handlers: Handler{
				Latency:     Latency{P75: "1s"},
				Reliability: Reliability{ErrorPercent: "0.01"},
				Throughput:  Throughput{RPM: "10k"},
			},
		}
		src := Defaults{
			Handlers: Handler{
				Latency:     Latency{P95: "2s"},
				Reliability: Reliability{ErrorPercent: "0.001"},
				Throughput:  Throughput{RPM: "100k"},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Defaults{
			Handlers: Handler{
				Latency:     Latency{P75: "1s", P95: "2s"},
				Reliability: Reliability{ErrorPercent: "0.001"},
				Throughput:  Throughput{RPM: "100k"},
			},
		}, dst)
	})
}

func TestHandler_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Handler
		assert.NotPanics(t, func() { dst.Merge(Handler{Name: "handler", Type: "brief"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Handler{Name: "handler-a", Type: "brief"}
		assert.NotPanics(t, func() { dst.Merge(Handler{Name: "handler-a", Type: "core", Latency: Latency{P99: "1s"}}) })
		assert.Empty(t, dst.Latency)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Handler{
			Name:        "handler",
			Type:        "brief",
			Latency:     Latency{P75: "1s"},
			Reliability: Reliability{ErrorPercent: "0.01"},
			Throughput:  Throughput{RPM: "10k"},
		}
		src := Handler{
			Name:        "handler",
			Type:        "brief",
			Latency:     Latency{P95: "2s"},
			Reliability: Reliability{ErrorCodes: []uint{499}},
			Throughput:  Throughput{RPM: "100k"},
		}

		dst.Merge(src)
		assert.Equal(t, Handler{
			Name: "handler",
			Type: "brief",
			Latency: Latency{
				P75: "1s",
				P95: "2s",
			},
			Reliability: Reliability{
				ErrorPercent: "0.01",
				ErrorCodes:   []uint{499},
			},
			Throughput: Throughput{RPM: "100k"},
		}, dst)
	})
}

func TestLatency_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Latency
		assert.NotPanics(t, func() { dst.Merge(Latency{P75: "10ms"}) })
		assert.Nil(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Latency{
			P75:  "10ms",
			P95:  "100ms",
			P98:  "150ms",
			P99:  "250ms",
			P999: "500ms",
		}
		src := Latency{
			P75:  "20ms",
			P95:  "200ms",
			P98:  "250ms",
			P99:  "350ms",
			P999: "600ms",
		}

		dst.Merge(src)
		assert.Equal(t, Latency{
			P75:  "20ms",
			P95:  "200ms",
			P98:  "250ms",
			P99:  "350ms",
			P999: "600ms",
		}, dst)
	})
}

func TestReliability_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Reliability
		assert.NotPanics(t, func() { dst.Merge(Reliability{ErrorPercent: "0.01"}) })
		assert.Nil(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Reliability{
			ErrorCodes:   []uint{499, 500},
			ErrorPercent: "0.01",
		}
		src := Reliability{
			ErrorCodes:   []uint{500, 503},
			ErrorPercent: "0.001",
		}

		dst.Merge(src)
		assert.Equal(t, Reliability{
			ErrorCodes:   []uint{499, 500, 503},
			ErrorPercent: "0.001",
		}, dst)
	})
}

func TestThroughput_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Throughput
		assert.NotPanics(t, func() { dst.Merge(Throughput{RPM: "1k"}) })
		assert.Nil(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Throughput{
			RPM: "60k",
			RPS: "1k",
		}
		src := Throughput{
			RPM: "120k",
			RPS: "2k",
		}

		dst.Merge(src)
		assert.Equal(t, Throughput{
			RPM: "120k",
			RPS: "2k",
		}, dst)
	})
}

func TestToken_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Token
		assert.NotPanics(t, func() { dst.Merge(Token{ID: "token"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Token{ID: "token-a"}
		assert.NotPanics(t, func() { dst.Merge(Token{ID: "token-b", Consumer: "client"}) })
		assert.Empty(t, dst.Consumer)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Token{
			ID:          "token",
			Consumer:    "client-a",
			Description: "Token description.",
			Engine:      "php",
			RequestedBy: "robot",
			Handlers: []Quota{
				{
					Name:        "quota-b",
					Type:        "rest",
					Scope:       "delete",
					Description: "Delete item.",
				},
				{
					Name:        "quota-a",
					Type:        "brief",
					Scope:       "fetch",
					Description: "Fetch item.",
				},
			},
		}
		src := Token{
			ID:          "token",
			Consumer:    "client-b",
			Description: "Token detailed description.",
			Engine:      "go",
			RequestedBy: "user",
			Handlers: []Quota{
				{
					Name:        "quota-a",
					Type:        "brief",
					Scope:       "fetch",
					Description: "Fetch item to calculation.",
				},
				{
					Name:        "quota-b",
					Type:        "core",
					Scope:       "update",
					Description: "Update item fields.",
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Token{
			ID:          "token",
			Consumer:    "client-b",
			Description: "Token detailed description.",
			Engine:      "go",
			RequestedBy: "user",
			Handlers: []Quota{
				{
					Name:        "quota-b",
					Type:        "rest",
					Scope:       "delete",
					Description: "Delete item.",
				},
				{
					Name:        "quota-a",
					Type:        "brief",
					Scope:       "fetch",
					Description: "Fetch item to calculation.",
				},
				{
					Name:        "quota-b",
					Type:        "core",
					Scope:       "update",
					Description: "Update item fields.",
				},
			},
		}, dst)
	})
}

func TestQuota_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Quota
		assert.NotPanics(t, func() { dst.Merge(Quota{Name: "quota"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Quota{Name: "quota"}
		assert.NotPanics(t, func() { dst.Merge(Quota{Name: "quota", Type: "brief"}) })
		assert.Empty(t, dst.Type)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Quota{
			Name:        "quota",
			Type:        "brief",
			Scope:       "fetch",
			Description: "Quota description.",
			Metadata:    nil,
		}
		src := Quota{
			Name:        "quota",
			Type:        "brief",
			Scope:       "fetch",
			Description: "Quota detailed description.",
			Metadata:    map[string]interface{}{"key": "value"},
		}

		dst.Merge(src)
		assert.Equal(t, Quota{
			Name:        "quota",
			Type:        "brief",
			Scope:       "fetch",
			Description: "Quota detailed description.",
			Metadata:    map[string]interface{}{"key": "value"},
		}, dst)
	})
}
