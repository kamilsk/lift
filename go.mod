module github.com/kamilsk/lift

go 1.11

require (
	github.com/kamilsk/platform v0.20.3
	github.com/pelletier/go-toml v1.6.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	go.octolab.org v0.0.7 // indirect
	go.octolab.org/toolkit/cli v0.0.6
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch
