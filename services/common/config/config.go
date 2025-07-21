package config

import (
	"log"
	"os"
	"path"

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

var DefaultPaths = newPathConf()

type pathConf struct {
	ConfigFile string
	DataFile   string
}

func newPathConf() *pathConf {
	pc := &pathConf{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dataFolder := path.Join(homeDir, ".local/share/movie-hub")
	configFolder := path.Join(homeDir, ".config/movie-hub")

	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		err = os.Mkdir(dataFolder, os.FileMode(0o744))
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		err = os.Mkdir(configFolder, os.FileMode(0o744))
		if err != nil {
			log.Fatal(err)
		}
	}

	pc.DataFile = path.Join(dataFolder, "data.json")
	pc.ConfigFile = path.Join(configFolder, "config.toml")

	return pc
}
