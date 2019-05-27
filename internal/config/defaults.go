package config

const (
	storagePostgres = "postgresql"
	storageMongo    = "mongodb"
	storageRabbit   = "rabbitmq"
	storageRedis    = "redis"
	storageSphinx   = "sphinx"
)

var defaults = Dependencies{
	{
		Name:    storagePostgres,
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
		Name:    storageMongo,
		Forward: []string{envMongoDSN},
		vars: map[string]string{
			envMongoDSN: "mongodb://localhost:27017",
		},
	},
	{
		Name:    storageRabbit,
		Forward: []string{envRabbitMQMaster, envRabbitMQBackup},
		vars: map[string]string{
			envRabbitMQMaster: "localhost:5672",
			envRabbitMQBackup: "localhost:5672",
		},
	},
	{
		Name:    storageRedis,
		Forward: []string{envRedisPort},
		vars: map[string]string{
			envRedisHost: "localhost",
			envRedisPort: "6379",
		},
	},
	{
		Name:    storageSphinx,
		Forward: []string{envSphinxPort},
		vars: map[string]string{
			envSphinxHost: "localhost",
			envSphinxPort: "9306",
		},
	},
}
