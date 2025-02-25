package main

import (
	"context"
	"flag"
	"fmt"
	test_task "github.com/EtoNeAnanasbI95/test-task"
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"github.com/EtoNeAnanasbI95/test-task/internal/handler"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"github.com/EtoNeAnanasbI95/test-task/internal/service"
	"github.com/EtoNeAnanasbI95/test-task/internal/storage"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

// @title Song Library API
// @version 1.0
// @description REST API для онлайн-библиотеки песен
// @host localhost:8080
// @BasePath /api
func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	envFilePath := filepath.Join(projectRoot, ".env")
	var flagEnvFilePath string
	flag.StringVar(&flagEnvFilePath, "config", envFilePath, "path to config file")
	flag.Parse()
	log.Println(flagEnvFilePath)
	cfg := config.MustLoadConfig(flagEnvFilePath)
	log.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting api",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.ApiPort))

	db := storage.MustInitDB(cfg.ConnectionString)

	log.Debug("Init db connection",
		slog.String("dsn", cfg.ConnectionString))

	r := repository.NewRepository(log, db)

	s := service.NewService(log, r)
	// TODO: добавить зависимость с клиентом open api
	apiHandler := handler.NewHandler(log, s)
	api := apiHandler.InitRouts()
	srv := new(test_task.Server)

	go func() {
		if err := srv.Run(api, cfg); err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("Api is running")
	if cfg.Env == envLocal {
		log.Info("Running in local mode",
			slog.String("URL", fmt.Sprintf("http://localhost:%d/swagger/index.html", cfg.ApiPort)))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("Shutting down server...")
	if err := srv.Stop(context.Background()); err != nil {
		log.Error(err.Error())
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
