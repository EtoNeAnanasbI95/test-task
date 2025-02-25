package repository

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

// TODO create repository interface
type Repository struct {
	log *slog.Logger
}

func NewRepository(l *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{log: l}
}
