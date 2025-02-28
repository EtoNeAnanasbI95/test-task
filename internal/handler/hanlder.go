package handler

import (
	"github.com/EtoNeAnanasbI95/test-task/internal/service"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log/slog"

	_ "github.com/EtoNeAnanasbI95/test-task/docs"
)

type Handler struct {
	log      *slog.Logger
	services *service.Service
}

func NewHandler(log *slog.Logger, s *service.Service) *Handler {
	return &Handler{
		log:      log,
		services: s,
	}
}

func (h *Handler) InitRouts(l *slog.Logger) *gin.Engine {
	router := gin.New()
	router.Use(h.CORSMiddleware)
	router.Use(sloggin.New(l))
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{
		api.GET("/songs", h.GetSongs())
		api.GET("/songs/:id/lyrics", h.GetSongLyrics())
		api.DELETE("/songs/:id", h.DeleteSong())
		api.PUT("/songs/:id", h.UpdateSong())
		api.POST("/songs", h.AddSong())
	}
	return router
}
