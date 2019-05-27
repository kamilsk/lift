package config

const (
	postgresStorage = "postgresql"
	mongoStorage    = "mongodb"
	rabbitStorage   = "rabbitmq"
	redisStorage    = "redis"
	sphinxStorage   = "sphinx"
)

var defaults = Dependencies{
	{
		Name:    postgresStorage,
		Forward: []string{envPgPort},
		vars: map[string]string{
			envPgHost:     "localhost",
			envPgPort:     "5432",
			envPgUser:     "postgres",
			envPgPassword: "",
			envPgDatabase: "master",
		},
	},
	{
		Name:    mongoStorage,
		Forward: []string{envMongoDSN},
		vars: map[string]string{
			envMongoDSN: "mongodb://localhost:27017",
		},
	},
	{
		Name:    rabbitStorage,
		Forward: []string{envRabbitMQMaster, envRabbitMQBackup},
		vars: map[string]string{
			envRabbitMQMaster: "localhost:5672",
			envRabbitMQBackup: "localhost:5672",
		},
	},
	{
		Name:    redisStorage,
		Forward: []string{envRedisPort},
		vars: map[string]string{
			envRedisHost: "localhost",
			envRedisPort: "6379",
		},
	},
	{
		Name:    sphinxStorage,
		Forward: []string{envSphinxPort},
		vars: map[string]string{
			encSphinxHost: "localhost",
			envSphinxPort: "9306",
		},
	},
}
