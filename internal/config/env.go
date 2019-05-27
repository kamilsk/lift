package config

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
