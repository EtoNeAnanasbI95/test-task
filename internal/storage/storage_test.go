package storage

import (
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"testing"
)

func TestMustInitDB(t *testing.T) {
	cfg := config.MustLoadConfig("../../.env")
	MustInitDB(cfg.ConnectionString)
}
