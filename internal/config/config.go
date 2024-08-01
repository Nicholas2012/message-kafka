package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Listen      string
	DatabaseDSN string
	KafkaBroker string
	KafkaTopic  string
}

func New() Config {
	_ = godotenv.Load()

	return Config{
		Listen:      getEnv("LISTEN", ":8080"),
		DatabaseDSN: getEnv("DATABASE_DSN", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:29092"),
		KafkaTopic:  getEnv("KAFKA_TOPIC", "test"),
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
