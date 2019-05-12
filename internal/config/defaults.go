package config

var defaults = []Dependency{
	{
		Name: "postgresql",
		vars: map[string]string{
			"PGHOST":     "",
			"PGPORT":     "",
			"PGUSER":     "",
			"PGPASSWORD": "",
			"PGDATABASE": "",
		},
	},
	{
		Name: "mongodb",
		vars: map[string]string{
			"MONGO_DSN": "",
		},
	},
	{
		Name: "rabbitmq",
		vars: map[string]string{
			"RABBITMQ_MASTER": "",
			"RABBITMQ_BACKUP": "",
		},
	},
	{
		Name: "redis",
		vars: map[string]string{
			"REDIS_HOST": "",
			"REDIS_PORT": "",
		},
	},
	{
		Name: "sphinx",
		vars: map[string]string{
			"SPHINX_HOST": "",
			"SPHINX_PORT": "",
		},
	},
}
