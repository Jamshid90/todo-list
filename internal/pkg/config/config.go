package config

import (
	"os"
)

type Config struct {
	App         string
	Environment string
	LogLevel    string
	RPCPort     string
	Context     struct {
		Timeout string
	}
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		Sslmode  string
	}
	HttpServer struct {
		Host         string
		Port         string
		ReadTimeout  string
		WriteTimeout string
		IdleTimeout  string
	}
}

func New() (*Config, error) {
	var config Config
	config.App = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.RPCPort = getEnv("RPC_PORT", ":9001")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "todotask")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "")
	config.DB.Sslmode = getEnv("POSTGRES_SSLMODE", "disable")

	config.HttpServer.Host = getEnv("HTTP_SERVER_HOST", "localhost")
	config.HttpServer.Port = getEnv("HTTP_SERVER_PORT", ":9007")
	config.HttpServer.ReadTimeout = getEnv("HTTP_SERVER_READ_TIMEOUT", "10s")
	config.HttpServer.WriteTimeout = getEnv("HTTP_SERVER_WRITE_TIMEOUT", "10s")
	config.HttpServer.IdleTimeout = getEnv("HTTP_SERVER_IDLE_TIMEOUT", "120s")

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
