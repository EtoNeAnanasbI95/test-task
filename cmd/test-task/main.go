// @title Song Library API
// @version 1.0
// @description REST API для тестового задания в компанию Effective Mobile, имитирующая библиотеку песен
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"context"
	"flag"
	"fmt"
	testtask "github.com/EtoNeAnanasbI95/test-task"
	openApi "github.com/EtoNeAnanasbI95/test-task/api"
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"github.com/EtoNeAnanasbI95/test-task/internal/handler"
	"github.com/EtoNeAnanasbI95/test-task/internal/lib/logger/sl"
	"github.com/EtoNeAnanasbI95/test-task/internal/repository"
	"github.com/EtoNeAnanasbI95/test-task/internal/service"
	"github.com/EtoNeAnanasbI95/test-task/internal/storage"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
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

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	envFilePath := filepath.Join(projectRoot, ".env")
	var flagEnvFilePath string
	flag.StringVar(&flagEnvFilePath, "config", envFilePath, "path to config file")
	flag.Parse()
	cfg := config.MustLoadConfig(flagEnvFilePath)
	log.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting api",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.ApiPort))

	connectionString := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v sslmode=%v",
		cfg.HostDB,
		cfg.PortDB,
		cfg.NameDB,
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.SslMode,
	)

	db := storage.MustInitDB(cfg.DBMS, connectionString)
	log.Debug("Init db connection",
		slog.String("dsn", connectionString))

	r := repository.NewRepository(log, db)
	log.Info("Init repository layer")

	apiClient := mustInitAPIClient(log, cfg.ExternalApiUrlBase)
	log.Debug("Init API client",
		slog.String("base url", cfg.ExternalApiUrlBase))

	s := service.NewService(log, apiClient, r)
	log.Info("Init service layer")

	apiHandler := handler.NewHandler(log, s)
	log.Info("Init handler layer")
	api := apiHandler.InitRouts(log)

	log.Debug("Init routing scheme")
	srv := new(testtask.Server)

	go func() {
		if err := srv.Run(api, cfg); err != nil {
			log.Error("Can not start web server: ", err.Error())
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

func mustInitAPIClient(log *slog.Logger, baseURL string) *openApi.Client {
	apiClient, err := openApi.NewClient(baseURL)
	if err != nil {
		log.Error("Can not init api client", sl.Err(err))
		panic("failed to initialize API client")
	}
	return apiClient
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	zerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	switch env {
	case envLocal:
		logger = slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &zerologLogger}.NewZerologHandler())
	case envProd:
		logger = slog.New(slogzerolog.Option{Level: slog.LevelInfo, Logger: &zerologLogger}.NewZerologHandler())
	default:
		logger = slog.New(slogzerolog.Option{Level: slog.LevelInfo, Logger: &zerologLogger}.NewZerologHandler())
	}
	return logger.With("env", env)
}
