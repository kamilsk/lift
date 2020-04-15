package config

import (
	"bufio"
	"path"
	"path/filepath"
	"regexp"

	"github.com/spf13/afero"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"
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

var goModRegexp = regexp.MustCompile(`(?is)^\s*module\s+([\w\-.]+(?:/[\w\-_]+)*)\s*$`)

// EngineSpecific returns the engine specific environment.
func EngineSpecific(engine Engine) Environment {
	switch fs := afero.NewOsFs(); engine.Name {
	case engineGo:
		env := Environment{}

		pkg, is := GoPackage(WorkDir{fs, engine.WorkDir})
		if is {
			env[envGoPackage] = pkg
			env[envGoImport] = pkg
		}

		mod, is := GoModule(WorkDir{fs, engine.WorkDir})
		if is {
			env[envGoModule] = mod
			env[envGoImport] = mod
		}

		return env
	case engineStatic, enginePHP, enginePython:
	}
	return nil
}

// GoModule returns go module name based on go.mod file.
func GoModule(wd WorkDir) (string, bool) {
	name := filepath.Join(wd.Path, fileGoMod)
	if _, err := wd.FS.Stat(name); err != nil {
		return "", false
	}
	file, err := wd.FS.Open(name)
	if err != nil {
		return "", false
	}
	defer safe.Close(file, unsafe.Ignore)
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
func GoPackage(wd WorkDir) (string, bool) {
	if _, err := wd.FS.Stat(filepath.Join(wd.Path, fileGoDep)); err == nil {
		location, pkg := wd.Path, make([]string, 0, 4)
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

// WorkDir holds a filesystem interface with a path of a current working directory.
type WorkDir struct {
	FS   afero.Fs
	Path string
}
