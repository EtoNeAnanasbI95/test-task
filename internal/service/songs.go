package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/EtoNeAnanasbI95/test-task/api"
	"github.com/EtoNeAnanasbI95/test-task/internal/lib/logger/sl"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"log/slog"
)

type SongsService struct {
	log       *slog.Logger
	apiClient api.ClientInterface
	repo      repository.Songs
}

func NewSongsService(l *slog.Logger, apiClient api.ClientInterface, repo repository.Songs) *SongsService {
	return &SongsService{
		log:       l,
		apiClient: apiClient,
		repo:      repo,
	}
}

func (s *SongsService) Add(ctx context.Context, model *models.SongInput) (int, error) {
	panic("implement me")
	return 0, nil
}

func (s *SongsService) GetLyrics(ctx context.Context, id int) error {
	panic("implement me")
	return nil
}

func (s *SongsService) GetAll(ctx context.Context, filter *models.SongFilter) ([]models.Song, error) {
	const OP = "SongsService.GetAll"

	log := s.log.With(slog.String("OP", OP))
	log.Debug("Fetching songs", "filter", filter)

	if !models.ValidateSongDate(filter) {
		log.Error("Release date have incorrect format")
		return nil, fmt.Errorf("%s: %w", OP, errors.New("release_date field have incorrect format"))
	}

	songs, err := s.repo.GetSongs(ctx, filter)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Debug("Request cancelled by client")
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("Request timed out")
		} else {
			log.Error("Failed to fetch songs", sl.Err(err))
		}
		return nil, fmt.Errorf("%s: %w", OP, err)
	}

	log.Debug("Songs count", "count", len(songs))
	return songs, nil
}

func (s *SongsService) Update(ctx context.Context, id int, model *models.SongUpdateInput) error {
	const OP = "SongsService.Update"

	log := s.log.With(
		slog.String("OP", OP),
		slog.Int("Updating song id", id),
	)
	log.Debug("Updating song", "song input", model)

	err := s.repo.UpdateSong(ctx, id, model)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Debug("Request cancelled by client")
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("Request timed out")
		} else {
			log.Error("Failed to update song", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", OP, err)
	}

	log.Debug("Successfully update song")
	return nil
}

func (s *SongsService) Delete(ctx context.Context, id int) error {
	const OP = "SongsService.Delete"

	log := s.log.With(
		slog.String("OP", OP),
		slog.Int("Deleting song id", id),
	)
	log.Debug("Deleting song")

	err := s.repo.DeleteSong(ctx, id)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Debug("Request cancelled by client")
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("Request timed out")
		} else {
			log.Error("Failed to delete song", sl.Err(err))
		}
		return fmt.Errorf("%s: %w", OP, err)
	}

	log.Debug("Successfully delete song")
	return nil
}
