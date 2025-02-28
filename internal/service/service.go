package service

import (
	"context"
	"github.com/EtoNeAnanasbI95/test-task/api"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"log/slog"
)

type Songs interface {
	Add(ctx context.Context, model *models.SongInput) (int, error)
	GetLyrics(ctx context.Context, id int) error
	GetAll(ctx context.Context, filter *models.SongFilter) ([]models.Song, error)
	Update(ctx context.Context, id int, model *models.SongUpdateInput) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	Songs
}

func NewService(l *slog.Logger, apiClient api.ClientInterface, r *repository.Repository) *Service {
	return &Service{
		Songs: NewSongsService(l, apiClient, r),
	}
}
