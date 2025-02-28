package handler

import (
	"github.com/EtoNeAnanasbI95/test-task/internal/lib/logger/sl"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) GetSongs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		const OP = "SongsHandler.GetSongs"
		log := h.log.With(slog.String("OP", OP))

		var filter *models.SongFilter
		if err := c.ShouldBindQuery(&filter); err != nil {
			log.Error("Invalid query params", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		songs, err := h.services.Songs.GetAll(ctx, filter)
		if err != nil {
			log.Error("Failed to fetch songs", sl.Err(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch songs"})
			return
		}

		log.Info("Successfully fetched songs")
		c.JSON(http.StatusOK, songs)
	}
}

func (h *Handler) GetSongLyrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic("implement me")
	}
}

func (h *Handler) DeleteSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		const OP = "SongsHandler.DeleteSong"
		log := h.log.With(slog.String("OP", OP))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("Invalid id query param", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id query param"})
			return
		}

		if err := h.services.Songs.Delete(ctx, id); err != nil {
			log.Error("Failed to delete song", sl.Err(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
			return
		}

		log.Info("Successfully delete song")
		c.Status(http.StatusNoContent)
	}
}

func (h *Handler) UpdateSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		const OP = "SongsHandler.UpdateSong"
		log := h.log.With(slog.String("OP", OP))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("Invalid id query param", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id query param"})
			return
		}

		var inputData *models.SongUpdateInput
		if err := c.ShouldBind(&inputData); err != nil {
			log.Error("Invalid query params", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		if err := h.services.Songs.Update(ctx, id, inputData); err != nil {
			log.Error("Failed to update song", sl.Err(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
			return
		}

		log.Info("Successfully update song")
		c.Status(http.StatusNoContent)
	}
}

func (h *Handler) AddSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic("implement me")
	}
}
