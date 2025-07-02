package config

import (
	"os"

	"github.com/joho/godotenv"
)

var EnvConfig = newEnvConfig()

type db struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

type envConfig struct {
	DataGrpcPort string
	DB           db
}

func newEnvConfig() *envConfig {
	godotenv.Load()
	return &envConfig{
		DataGrpcPort: os.Getenv("DATA_GRPC_PORT"),
		DB: db{
			Host: os.Getenv("POSTGRES_HOST"),
			Port: os.Getenv("POSTGRES_PORT"),
			Name: os.Getenv("POSTGRES_DB_NAME"),
			User: os.Getenv("POSTGRES_USER"),
			Pass: os.Getenv("POSTGRES_PASS"),
		},
	}
}
