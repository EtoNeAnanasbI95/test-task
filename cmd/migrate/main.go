package main

import (
	"errors"
	"flag"
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	envFilePath := filepath.Join(projectRoot, ".env")
	migrationPath := filepath.Join(projectRoot, "migrations")
	var flagEnvFilePath string
	var flagMigrationPath string
	flag.StringVar(&flagEnvFilePath, "config", envFilePath, "path to config file")
	flag.StringVar(&flagMigrationPath, "migrations", migrationPath, "path to migrations directory")
	flag.Parse()
	log.Println(flagEnvFilePath)
	log.Println(flagMigrationPath)
	cfg := config.MustLoadConfig(flagEnvFilePath)
	log.Println(cfg)

	if cfg.ConnectionString == "" {
		panic("connection string is required")
	}

	m, err := migrate.New(
		"file://"+flagMigrationPath,
		cfg.ConnectionString,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply")
			return
		}
		panic(err)
	}

	log.Println("migrations applied")
}
