package config

import (
	"os"

	"github.com/joho/godotenv"
)

var EnvConfig = newEnvConfig()

type envConfig struct {
	DataGrpcPort string
}

func newEnvConfig() *envConfig {
	godotenv.Load()
	return &envConfig{
		DataGrpcPort: os.Getenv("DATA_GRPC_PORT"),
	}
}
