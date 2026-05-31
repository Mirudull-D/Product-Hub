package config

import "os"

var Envs = initConfig()

func initConfig() Config {
	return Config{
		ConnString: getEnv("connString",
			"postgres://mirudull:1234567890@localhost:5432/goapplication?sslmode=disable"),
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
