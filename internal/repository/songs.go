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

func (r *SongsRepository) AddSong(ctx context.Context, model *models.SongUpdateInput) (int, error) {
	const OP = "SongsRepository.AddSong"
	log := r.log.With(slog.String("OP", OP))

	log.Debug("Add new song", "new song", model)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error("Failed to start transaction", sl.Err(err))
		return 0, fmt.Errorf("%s: failed to start transaction: %w", OP, err)
	}

	query := `
        INSERT INTO "Songs" ("group", "song", "release_date", "text", "link")
        VALUES ($1, $2, $3, $4, $5)
        RETURNING "id"`

	var id int
	err = tx.QueryRowxContext(ctx, query,
		model.Group,
		model.Song,
		model.ReleaseDate,
		model.Text,
		model.Link,
	).Scan(&id)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Error("Failed to rollback transaction", "rollback_error", rollbackErr, sl.Err(err))
			return 0, fmt.Errorf("%s: failed to rollback after error: %w (original error: %w)", OP, rollbackErr, err)
		}
		log.Error("Failed to insert song", "query", query, sl.Err(err))
		return 0, fmt.Errorf("%s: failed to insert song: %w", OP, err)
	}

	if err := tx.Commit(); err != nil {
		log.Error("Failed to commit transaction", sl.Err(err))
		return 0, fmt.Errorf("%s: failed to commit transaction: %w", OP, err)
	}

	log.Debug("Successfully inserted new song", "id", id)
	return id, nil
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

	if filter.Limit != 0 {
		query += ` LIMIT ?`
		args = append(args, filter.Limit)
	}
	if filter.Offset != 0 {
		query += ` OFFSET ?`
		args = append(args, filter.Offset)
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

	query := fmt.Sprintf(`DELETE FROM "%v" WHERE "id" = ?;`, songsTable)
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Error("Failed to delete song", "query", query, sl.Err(err))
		return fmt.Errorf("%s: %w", OP, err)
	}
	log.Debug("Successfully delete song")
	return nil
}
