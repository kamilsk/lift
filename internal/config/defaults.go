package config

const (
	pgStorage     = "postgresql"
	mongoStorage  = "mongodb"
	rabbitStorage = "rabbitmq"
	redisStorage  = "redis"
	sphinxStorage = "sphinx"
)

var defaults = Dependencies{
	{
		Name: pgStorage,
		vars: map[string]string{
			"PGHOST":     "localhost",
			"PGPORT":     "5432",
			"PGUSER":     "postgres",
			"PGPASSWORD": "",
			"PGDATABASE": "postgres",
		},
	},
	{
		Name: mongoStorage,
		vars: map[string]string{
			"MONGO_DSN": "mongodb://localhost:27017",
		},
	},
	{
		Name: rabbitStorage,
		vars: map[string]string{
			"RABBITMQ_MASTER": "localhost:5672",
			"RABBITMQ_BACKUP": "localhost:5672",
		},
	},
	{
		Name: redisStorage,
		vars: map[string]string{
			"REDIS_HOST": "localhost",
			"REDIS_PORT": "6379",
		},
	},
	{
		Name: sphinxStorage,
		vars: map[string]string{
			"SPHINX_HOST": "localhost",
			"SPHINX_PORT": "9306",
		},
	},
}
