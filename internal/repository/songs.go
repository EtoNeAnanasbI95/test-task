package repository

import (
	"context"
	"fmt"
	"github.com/EtoNeAnanasbI95/test-task/internal/lib/logger/sl"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type SongsRepository struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewSongsRepository(l *slog.Logger, db *sqlx.DB) *SongsRepository {
	return &SongsRepository{
		log: l,
		db:  db,
	}
}

func (r *SongsRepository) GetSongs(ctx context.Context, filter *models.SongFilter) ([]models.Song, error) {
	const OP = "SongsRepository.GetSongs"
	log := r.log.With(slog.String("OP", OP))

	query := fmt.Sprintf(`SELECT * FROM "%v" WHERE 1=1`, songsTable)
	args := []any{}

	if filter.Group != "" {
		query += ` AND "group" ILIKE ?`
		args = append(args, "%"+filter.Group+"%")
	}
	if filter.Link != "" {
		query += ` AND "link" ILIKE ?`
		args = append(args, "%"+filter.Link+"%")
	}
	if filter.Song != "" {
		query += ` AND "song" ILIKE ?`
		args = append(args, "%"+filter.Song+"%")
	}
	if filter.Text != "" {
		query += ` AND "text" ILIKE ?`
		args = append(args, "%"+filter.Text+"%")
	}
	if filter.ReleaseDate != "" {
		query += ` AND "release_date" ILIKE ?`
		args = append(args, "%"+filter.ReleaseDate+"%")
	}

	// TODO: доделать пагинацию

	if filter.Page != 1 {
	}
	if filter.PageSize != 1 {
	}

	query, args, err := sqlx.In(query, args...)
	log.Debug("Created query", "query", query)
	if err != nil {
		log.Error("Arguments parse error", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", OP, err)
	}
	query = r.db.Rebind(query)

	var songs []models.Song
	err = r.db.SelectContext(ctx, &songs, query, args...)
	if err != nil {
		log.Error("Failed to fetch songs", "query", query, sl.Err(err))
		return nil, fmt.Errorf("%s: %w", OP, err)
	}
	log.Debug("Songs fetched")
	return songs, nil
}

func (r *SongsRepository) UpdateSong(ctx context.Context, id int, model *models.SongUpdateInput) error {
	const OP = "SongsRepository.UpdateSong"
	log := r.log.With(slog.String("OP", OP))

	query := fmt.Sprintf(`UPDATE "%v" SET`, songsTable)
	args := []any{}

	if model.Group != "" {
		query += ` "group" = ?,`
		args = append(args, model.Group)
	}
	if model.Link != "" {
		query += ` "link" = ?,`
		args = append(args, model.Link)
	}
	if model.Song != "" {
		query += ` "song" = ?,`
		args = append(args, model.Song)
	}
	if model.Text != "" {
		query += ` "text" = ?,`
		args = append(args, model.Text)
	}
	if model.ReleaseDate != "" {
		query += ` "release_date" = ?`
		args = append(args, model.ReleaseDate)
	}

	query += ` WHERE "id" = ?`
	args = append(args, id)

	query, args, err := sqlx.In(query, args...)
	log.Debug("Created query", "query", query)
	if err != nil {
		log.Error("Arguments parse error", sl.Err(err))
		return fmt.Errorf("%s: %w", OP, err)
	}
	query = r.db.Rebind(query)

	var songs []models.Song
	err = r.db.SelectContext(ctx, &songs, query, args...)
	if err != nil {
		log.Error("Failed to update song", "query", query, sl.Err(err))
		return fmt.Errorf("%s: %w", OP, err)
	}
	log.Debug("Song updated")
	return nil
}

func (r *SongsRepository) DeleteSong(ctx context.Context, id int) error {
	const OP = "SongsRepository.DeleteSong"
	log := r.log.With(slog.String("OP", OP))

	query := fmt.Sprintf(`DELETE FROM "%v" WHERE id = ?;`, songsTable)
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Error("Failed to delete song", "query", query, sl.Err(err))
		return fmt.Errorf("%s: %w", OP, err)
	}
	log.Debug("Successfully delete song")
	return nil
}
