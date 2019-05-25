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
		Name:    pgStorage,
		Forward: []string{"PGPORT"},
		vars: map[string]string{
			"PGHOST":     "localhost",
			"PGPORT":     "5432",
			"PGUSER":     "postgres",
			"PGPASSWORD": "",
			"PGDATABASE": "master",
		},
	},
	{
		Name:    mongoStorage,
		Forward: []string{"MONGO_DSN"},
		vars: map[string]string{
			"MONGO_DSN": "mongodb://localhost:27017",
		},
	},
	{
		Name:    rabbitStorage,
		Forward: []string{"RABBITMQ_MASTER", "RABBITMQ_BACKUP"},
		vars: map[string]string{
			"RABBITMQ_MASTER": "localhost:5672",
			"RABBITMQ_BACKUP": "localhost:5672",
		},
	},
	{
		Name:    redisStorage,
		Forward: []string{"REDIS_PORT"},
		vars: map[string]string{
			"REDIS_HOST": "localhost",
			"REDIS_PORT": "6379",
		},
	},
	{
		Name:    sphinxStorage,
		Forward: []string{"SPHINX_PORT"},
		vars: map[string]string{
			"SPHINX_HOST": "localhost",
			"SPHINX_PORT": "9306",
		},
	},
}
