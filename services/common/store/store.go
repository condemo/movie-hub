package store

import "github.com/jmoiron/sqlx"

var tables = `
CREATE TABLE IF NOT EXISTS media (
	id integer PRIMARY KEY,
	type varchar(15) NOT NULL,
	title varchar(30) NOT NULL,
	year smallint NOT NULL,
	genres text NOT NULL,
	seasons smallint NULL,
	caps smallint NULL,
	description text NOT NULL,
	rating smallint NOT NULL,
	image text,
	fav boolean DEFAULT false,
	viewed boolean DEFAULT false,
	UNIQUE(title)
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
