package service

import (
	"context"
	"encoding/json"
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
	const OP = "SongsService.Add"

	log := s.log.With(slog.String("OP", OP))
	log.Debug("Add song", "model", model)

	rawData, err := s.apiClient.GetInfo(ctx, &api.GetInfoParams{
		Group: model.Group,
		Song:  model.Song,
	})
	if err != nil {
		log.Error("Can not get full song data out external api", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", OP, errors.New("can not get full song data"))
	}

	defer rawData.Body.Close()

	var songDetails api.SongDetail
	if err := json.NewDecoder(rawData.Body).Decode(&songDetails); err != nil {
		log.Error("Failed to decode response body", sl.Err(err))
		return 0, fmt.Errorf("%s: failed to decode song data: %w", OP, err)
	}

	id, err := s.repo.AddSong(ctx, &models.SongUpdateInput{
		Group:       model.Group,
		Song:        model.Song,
		ReleaseDate: songDetails.ReleaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	})

	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Debug("Request cancelled by client")
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Warn("Request timed out")
		} else {
			log.Error("Failed to fetch songs", sl.Err(err))
		}
		return 0, fmt.Errorf("%s: %w", OP, err)
	}

	log.Debug("Added song id", "id", id)
	return id, nil
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
