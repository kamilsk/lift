package cnf

import "go.octolab.org/toolkit/config"

// Features defines a list of available features.
var Features = config.Features{
	{
		Name:    "paas",
		Enabled: true,
	},
}
