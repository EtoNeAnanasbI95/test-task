package service

import (
	"bytes"
	"context"
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"github.com/EtoNeAnanasbI95/test-task/internal/service/mocks"
	"github.com/EtoNeAnanasbI95/test-task/internal/storage"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"log/slog"
	"net/http"
	"os"
	"testing"
)

func TestSongsService_Add(t *testing.T) {
	cfg := config.MustLoadConfig("../../.env")
	db := storage.MustInitDB(cfg.ConnectionString)
	zerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	log := slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &zerologLogger}.NewZerologHandler())
	repo := repository.NewSongsRepository(log, db)
	mockApi := mocks.NewClientInterface(t)
	model := &models.SongInput{
		Group: "Muse",
		Song:  "Supermassive Black Hole",
	}
	mockApi.On("GetInfo", mock.Anything, mock.Anything, mock.Anything).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"group": "Muse", "song": "Supermassive Black Hole", "releaseDate": "2006-07-16", "text": "Ooh baby...", "link": "https://youtube.com/..."}`)),
	}, nil)

	songService := NewSongsService(log, mockApi, repo)

	id, err := songService.Add(context.Background(), model)
	assert.NoError(t, err, "expect no error")
	assert.NotEmpty(t, id)
}
