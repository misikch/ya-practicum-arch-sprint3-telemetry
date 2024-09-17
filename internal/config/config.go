package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServicePort string `envconfig:"SERVICE_PORT" default:"8081"`
	PgHost      string `envconfig:"PGHOST" required:"true"`
	PgPort      string `envconfig:"PGPORT" required:"true"`
	PgUser      string `envconfig:"PGUSER" required:"true"`
	PgPass      string `envconfig:"PGPASSWORD" required:"true"`
	PgDatabase  string `envconfig:"PGDATABASE" required:"true"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("deviceManager", &cfg)
	return cfg, err
}
