package paas_test

import (
	"strings"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestNFR(t *testing.T) {
	data := `
[default.handlers]
  error_codes = [503, 504]

[[handlers]]
  name = "countWords"
  type = "brief"
  max_rpm = "10k"
  errors_percent = "0.3"
  latency_p99 = "100ms"

[[handlers]]
  name = "/createTable"
  type = "core"
  max_rpm = "1000"
  latency_p999 = "1s"

[[handlers]]
  name = "/drop/table/[a-z0-9-]+$"
  type = "rest"
  error_codes = [200, 500]
  max_rpm = "100"

[[quota]]
  id = "client"
  consumer = "https://www.example.com"
  engine = "php"
  requested_by = "https://en.wikipedia.org/wiki/Robot"
  description = """
A robot may not injure a human being or, through inaction, allow a human being to come to harm.
A robot must obey the orders given it by human beings except where such orders would conflict with the First Law.
A robot must protect its own existence as long as such protection does not conflict with the First or Second Law.
"""

  [[quota.handlers]]
    name = "countWords"
    scope = "read"
    latency_p99 = "100ms"
    max_rpm = "1k"

    [quota.handlers.metadata]
      critical = true
      fields = ["a", "b", "c"]
      features = ["deadline"]
`

	var nfr NFR
	assert.NoError(t, toml.NewDecoder(strings.NewReader(data)).Decode(&nfr))
	assert.Equal(t, NFR{
		Defaults: Defaults{
			Handlers: Handler{
				Reliability: Reliability{
					ErrorCodes: []uint{503, 504},
				},
			},
		},
		Handlers: []Handler{
			{
				Name: "countWords",
				Type: "brief",
				Latency: Latency{
					P99: "100ms",
				},
				Reliability: Reliability{
					ErrorPercent: "0.3",
				},
				Throughput: Throughput{
					RPM: "10k",
				},
			},
			{
				Name: "/createTable",
				Type: "core",
				Latency: Latency{
					P999: "1s",
				},
				Throughput: Throughput{
					RPM: "1000",
				},
			},
			{
				Name: "/drop/table/[a-z0-9-]+$",
				Type: "rest",
				Reliability: Reliability{
					ErrorCodes: []uint{200, 500},
				},
				Throughput: Throughput{
					RPM: "100",
				},
			},
		},
		Quota: []Token{
			{
				ID:          "client",
				Consumer:    "https://www.example.com",
				Engine:      "php",
				RequestedBy: "https://en.wikipedia.org/wiki/Robot",
				Description: strings.TrimLeft(`
A robot may not injure a human being or, through inaction, allow a human being to come to harm.
A robot must obey the orders given it by human beings except where such orders would conflict with the First Law.
A robot must protect its own existence as long as such protection does not conflict with the First or Second Law.
`, "\n"),
				Handlers: []Quota{
					{
						Name:  "countWords",
						Scope: "read",
						Latency: Latency{
							P99: "100ms",
						},
						Throughput: Throughput{
							RPM: "1k",
						},
						Metadata: map[string]interface{}{
							"critical": true,
							"fields":   []interface{}{"a", "b", "c"},
							"features": []interface{}{"deadline"},
						},
					},
				},
			},
		},
	}, nfr)
}
