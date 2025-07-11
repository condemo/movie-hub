package store

import (
	"context"

	"github.com/condemo/movie-hub/services/common/types"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetLastUpdates(context.Context) ([]*types.MediaResume, error)
	GetOneMedia(context.Context, int64) (*types.Media, error)
	InsertMedia(context.Context, *types.Media) error
	InsertBulkMedia(context.Context, []types.Media) error
	DeleteMedia(context.Context, int64) error
	UpdateMedia(context.Context, *types.Media) error
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetLastUpdates(ctx context.Context) ([]*types.MediaResume, error) {
	mr := []*types.MediaResume{}
	err := s.db.SelectContext(ctx, &mr, `SELECT 
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

func (s *Storage) InsertMedia(ctx context.Context, m *types.Media) error {
	rows, err := s.db.NamedQueryContext(ctx, `INSERT INTO media
		(type, title, year,
		genres, seasons, caps, description, rating, image, fav, viewed)
		VALUES (:type, :title, :year,:genres, :seasons,:caps,:description,:rating,
		:image,:fav,:viewed) RETURNING *;`, m)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.StructScan(m); err != nil {
			return err
		}
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) InsertBulkMedia(ctx context.Context, m []types.Media) error {
	// TODO: recibir las rows de `NamedQuery`
	_, err := s.db.NamedExecContext(ctx, `INSERT INTO media (type, title, year,
		genres, seasons, caps, description, rating, image, fav, viewed)
		VALUES (:type, :title, :year,:genres, :seasons,:caps,:description,:rating,
		:image,:fav,:viewed)`, m)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateMedia(ctx context.Context, m *types.Media) error {
	rows, err := s.db.NamedQueryContext(ctx, `UPDATE media SET 
		type=:type, title=:title, year=:year, genres=:genres, seasons=:seasons,
		caps=:caps, description=:description, rating=:rating, image=:image,
		fav=:fav, viewed=:viewed WHERE id=:id RETURNING *`, m)
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteMedia(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM media WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}
