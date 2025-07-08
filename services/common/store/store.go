package store

import (
	"context"

	"github.com/condemo/movie-hub/services/common/types"
	"github.com/jmoiron/sqlx"
)

// TODO:
type Store interface {
	GetLastUpdates() ([]*types.MediaResume, error)
	GetOneMedia(ctx context.Context, id int64) (*types.Media, error)
	InsertBulkMedia([]types.Media) error
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetLastUpdates() ([]*types.MediaResume, error) {
	mr := []*types.MediaResume{}
	err := s.db.Select(&mr, `SELECT 
		id, type, title,genres,description, image, fav, viewed
		FROM media ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}

	return mr, nil
}

func (s *Storage) GetOneMedia(ctx context.Context, id int64) (*types.Media, error) {
	movie := new(types.Media)
	err := s.db.Get(movie, "SELECT * FROM media WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Storage) InsertBulkMedia(m []types.Media) error {
	// TODO: recibir las rows de `NamedQuery`
	_, err := s.db.NamedExec(`INSERT INTO media (type, title, year,
		genres, seasons, caps, description, rating, image, fav, viewed)
		VALUES (:type, :title, :year,:genres, :seasons,:caps,:description,:rating,
		:image,:fav,:viewed)`, m)
	if err != nil {
		return err
	}

	return nil
}
