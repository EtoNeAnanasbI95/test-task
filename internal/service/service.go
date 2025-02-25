package service

import (
	"github.com/EtoNeAnanasbI95/test-task/api"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"log/slog"
)

type Songs interface {
	Create() (int, error)
	Get(id int) error
	GetAll() error
	Update(id int) error
	Delete(id int) error
}

type Service struct {
	log *slog.Logger
	Songs
}

func NewService(l *slog.Logger, apiClient api.ClientInterface, r *repository.Repository) *Service {
	return &Service{
		log: l,
		//Songs: ,
	}
}
