module github.com/kamilsk/lift

require (
	github.com/kamilsk/platform v0.18.0
	github.com/pelletier/go-toml v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	go.octolab.org/toolkit/cli v0.0.3
)

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch
