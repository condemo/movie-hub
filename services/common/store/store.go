package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/condemo/movie-hub/services/common/protogen/pb"
	"github.com/condemo/movie-hub/services/common/types"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetLastUpdates(context.Context, int32) ([]*types.MediaResume, error)
	GetOneMedia(context.Context, int64) (*types.Media, error)
	GetMediaFiltered(context.Context, *pb.MediaFilteredRequest) ([]*types.MediaResume, error)
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

func (s *Storage) GetLastUpdates(ctx context.Context, limit int32) ([]*types.MediaResume, error) {
	mr := []*types.MediaResume{}
	err := s.db.SelectContext(ctx, &mr, `SELECT 
		id, media_type, title, genres, description, image, fav, viewed
		FROM media ORDER BY id DESC LIMIT $1`, limit)
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

func (s *Storage) GetMediaFiltered(ctx context.Context, fb *pb.MediaFilteredRequest) ([]*types.MediaResume, error) {
	var sb strings.Builder
	mr := []*types.MediaResume{}
	sb.WriteString(fmt.Sprintf(`SELECT
		id, media_type, title, genres, description, image, fav,
		viewed FROM media WHERE %s=true`, fb.GetFilter().String()))
	if fb.GetLimit() > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", fb.GetLimit()))
	}

	err := s.db.SelectContext(ctx, &mr, sb.String())
	if err != nil {
		return nil, err
	}
	return mr, nil
}

func (s *Storage) InsertMedia(ctx context.Context, m *types.Media) error {
	rows, err := s.db.NamedQueryContext(ctx, `INSERT INTO media
		(media_type, title, release_year, first_air,
		genres, seasons, caps, description, rating, runtime, image, fav, viewed)
		VALUES (:media_type, :title, :release_year, :first_air,
		:genres, :seasons, :caps, :description, :rating, :runtime,
		:image,:fav,:viewed)
		 ON CONFLICT (title) DO NOTHING RETURNING *`, m)
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
	_, err := s.db.NamedExecContext(ctx, `INSERT INTO media (media_type, title,
		release_year,	first_air, genres, seasons, caps, description,
		rating, runtime, image, fav, viewed)
		VALUES (:media_type, :title, :release_year, :first_air, :genres,
		:seasons, :caps, :description, :rating, :runtime, :image, :fav, :viewed)
		ON CONFLICT (title) DO NOTHING`, m)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateMedia(ctx context.Context, m *types.Media) error {
	rows, err := s.db.NamedQueryContext(ctx, `UPDATE media SET 
		media_type=:media_type, title=:title, release_year=:release_year,
		first_air=:first_air, genres=:genres, seasons=:seasons,
		caps=:caps, description=:description, rating=:rating, runtime=:runtime,
		image=:image, fav=:fav, viewed=:viewed WHERE id=:id RETURNING *`, m)
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
