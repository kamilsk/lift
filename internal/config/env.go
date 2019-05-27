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
	encSphinxHost     = "SPHINX_HOST"
	envSphinxPort     = "SPHINX_PORT"
)

// EngineSpecific returns the engine specific environment.
func EngineSpecific(engine Engine) Environment {
	return nil
}
