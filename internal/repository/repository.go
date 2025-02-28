package repository

import (
	"context"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const songsTable = "Songs"

type Songs interface {
	GetSongs(ctx context.Context, filter *models.SongFilter) ([]models.Song, error)
	DeleteSong(ctx context.Context, id int) error
	UpdateSong(ctx context.Context, id int, model *models.SongUpdateInput) error
	AddSong(ctx context.Context, model *models.SongUpdateInput) (int, error)
}

type Repository struct {
	Songs
}

func NewRepository(l *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Songs: NewSongsRepository(l, db),
	}
}
