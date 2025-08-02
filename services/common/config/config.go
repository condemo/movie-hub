package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"time"

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
	pc.ConfigFile = path.Join(configFolder, "config.json")

	return pc
}

var General = newGeneralConf()

type generalConf struct {
	UpdateTimeInterval time.Duration
	DefaultDataLimit   int32
}

func newGeneralConf() *generalConf {
	gc := &generalConf{}
	if _, err := os.Stat(DefaultPaths.ConfigFile); os.IsNotExist(err) {
		f, err := os.Create(DefaultPaths.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		gc.UpdateTimeInterval = time.Duration(time.Hour * 24 * 3)
		gc.DefaultDataLimit = 20

		err = json.NewEncoder(f).Encode(gc)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err := os.Open(DefaultPaths.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = json.NewDecoder(f).Decode(gc)
		if err != nil {
			log.Fatal(err)
		}
	}

	return gc
}

// TODO: implementar un el método `Save` para guardar la config en un archivo
