package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		return Config{
			ConnString: "postgres://mirudull:1234567890@192.168.0.108:5432/goapplication?sslmode=disable",
			Port:       ":8080",
		}
	}
	return Config{
		ConnString: getEnv("connString",
			"postgres://mirudull:1234567890@192.168.0.108:5432/goapplication?sslmode=disable"),
		Port: getEnv("port", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Config struct {
	ConnString string
	Port       string
}
