package config

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/kamilsk/platform/pkg/safe"
)

const (
	envPgHost         = "PGHOST"
	envPgPort         = "PGPORT"
	envPgUser         = "PGUSER"
	envPgPassword     = "PGPASSWORD"
	envPgDatabase     = "PGDATABASE"
	envMongoDSN       = "MONGO_DSN"
	envRabbitMQMaster = "RABBITMQ_MASTER"
	envRabbitMQBackup = "RABBITMQ_BACKUP"
	envRedisHost      = "REDIS_HOST"
	envRedisPort      = "REDIS_PORT"
	envSphinxHost     = "SPHINX_HOST"
	envSphinxPort     = "SPHINX_PORT"
	envGoImport       = "GOIMPORT"
	envGoModule       = "GOMODULE"
	envGoPackage      = "GOPACKAGE"
)

const (
	engineGo     = "golang"
	engineStatic = "static"
	enginePHP    = "php"
	enginePython = "python-flask"
)

const (
	fileGoDep = "Gopkg.toml"
	fileGoMod = "go.mod"
)

var goModRegexp = regexp.MustCompile(`(?is)^\s*module\s+([\w./]+)\s*$`)

// EngineSpecific returns the engine specific environment.
func EngineSpecific(engine Engine) Environment {
	switch engine.Name {
	case engineGo:
		env := Environment{}
		return env
	case engineStatic, enginePHP, enginePython:
	}
	return nil
}

// GoModule returns go module name based on go.mod file.
func GoModule(location string) (string, bool) {
	name := filepath.Join(location, fileGoMod)
	if _, err := os.Stat(name); err != nil {
		return "", false
	}
	file, err := os.Open(name)
	if err != nil {
		return "", false
	}
	defer safe.Close(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if goModRegexp.MatchString(line) {
			return goModRegexp.FindStringSubmatch(line)[1], true
		}
	}
	return "", false
}

// GoPackage returns go package name based on its path.
func GoPackage(location string) (string, bool) {
	if _, err := os.Stat(filepath.Join(location, fileGoDep)); err == nil {
		pkg := make([]string, 0, 4)
		for {
			dir, file := filepath.Split(location)
			if dir == "" || file == "src" || len(pkg) == cap(pkg) {
				break
			}
			location = filepath.Dir(dir)
			pkg = append(pkg, file)
		}
		for left, right := 0, len(pkg)-1; left < right; left, right = left+1, right-1 {
			pkg[left], pkg[right] = pkg[right], pkg[left]
		}
		return path.Join(pkg...), true
	}
	return "", false
}
