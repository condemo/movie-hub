package store

import (
	"fmt"
	"log"

	"github.com/condemo/movie-hub/services/common/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresqlStorage() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		config.EnvConfig.DB.Host, config.EnvConfig.DB.User,
		config.EnvConfig.DB.Pass, config.EnvConfig.DB.Port,
		config.EnvConfig.DB.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
