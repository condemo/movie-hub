package store

import "github.com/jmoiron/sqlx"

var tables = `
CREATE TABLE media (
	id integer PRIMARY KEY,
	title varchar(30) NOT NULL,
	year smallint NOT NULL,
	genres text,
	seasons smallint NULL,
	caps smallint NULL,
	description text,
	rating smallint,
	image text,
	fav boolean,
	viewed boolean
	);
`

// TODO:
type Store interface{}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}
