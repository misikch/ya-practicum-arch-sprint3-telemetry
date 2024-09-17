package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServicePort   string `envconfig:"SERVICE_PORT" default:"8081"`
	MongoHost     string `envconfig:"MONGO_HOST" required:"true"`
	MongoPort     string `envconfig:"MONGO_PORT" required:"true"`
	MongoUser     string `envconfig:"MONGO_USER" required:"true"`
	MongoPass     string `envconfig:"MONGO_PASSWORD" required:"true"`
	MongoDatabase string `envconfig:"MONGO_DATABASE" required:"true"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("deviceManager", &cfg)
	return cfg, err
}
