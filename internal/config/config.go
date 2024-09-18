package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServicePort         string `envconfig:"SERVICE_PORT" default:"8082"`
	MongoHost           string `envconfig:"MONGO_HOST" required:"true"`
	MongoPort           string `envconfig:"MONGO_PORT" required:"true"`
	MongoUser           string `envconfig:"MONGO_USER" required:"true"`
	MongoPass           string `envconfig:"MONGO_PASSWORD" required:"true"`
	MongoDatabase       string `envconfig:"MONGO_DATABASE" required:"true"`
	KafkaBroker         string `envconfig:"KAFKA_BROKER" required:"true"` //example: KAFKA_BROKER=kafka-broker:9092 or KAFKA_BROKER=kafka-broker1:9092,kafka-broker2:9092,kafka-broker3:9092
	KafkaTelemetryTopic string `envconfig:"KAFKA_TELEMETRY_TOPIC" required:"true"`
	KafkaDevicesTopic   string `envconfig:"KAFKA_DEVICES_TOPIC" required:"true"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	err = envconfig.Process("deviceManager", cfg)
	return cfg, err
}
