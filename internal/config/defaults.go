package config

var defaults = []Dependency{
	{
		Name: "postgresql",
		vars: map[string]string{
			"PGHOST":     "localhost",
			"PGPORT":     "5432",
			"PGUSER":     "postgres",
			"PGPASSWORD": "",
			"PGDATABASE": "postgres",
		},
	},
	{
		Name: "mongodb",
		vars: map[string]string{
			"MONGO_DSN": "mongodb://localhost:27017",
		},
	},
	{
		Name: "rabbitmq",
		vars: map[string]string{
			"RABBITMQ_MASTER": "localhost:5672",
			"RABBITMQ_BACKUP": "localhost:5672",
		},
	},
	{
		Name: "redis",
		vars: map[string]string{
			"REDIS_HOST": "localhost",
			"REDIS_PORT": "6379",
		},
	},
	{
		Name: "sphinx",
		vars: map[string]string{
			"SPHINX_HOST": "localhost",
			"SPHINX_PORT": "9306",
		},
	},
}
