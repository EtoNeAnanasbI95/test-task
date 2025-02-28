package handler

import (
	"github.com/EtoNeAnanasbI95/test-task/internal/lib/logger/sl"
	"github.com/EtoNeAnanasbI95/test-task/pkg/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// GetSongs Получение списка песен
// @Summary Получение списка песен
// @Description Возвращает список песен с фильтрацией по всем полям и пагинацией
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Фильтр по группе"
// @Param song query string false "Фильтр по названию песни"
// @Param releaseDate query string false "Фильтр по дате релиза (формат: YYYY-MM-DD)"
// @Param text query string false "Фильтр по тексту песни"
// @Param link query string false "Фильтр по ссылке"
// @Param offset query int false "Смещение для пагинации" default(0)
// @Param limit query int false "Лимит записей" default(0)
// @Success 200 {array} models.Song "Список песен"
// @Failure 400 {object} map[string]string "Некорректные параметры запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs [get]
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

// GetSongLyrics Получение текста песни с пагинацией
// @Summary Получение текста песни с пагинацией
// @Description Возвращает текст песни, разделённый на куплеты, с пагинацией
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param offset query int false "Смещение для пагинации куплетов" default(0)
// @Param limit query int false "Лимит куплетов" default(0)
// @Success 200 {object} map[string][]string "Текст песни в виде массива куплетов"
// @Failure 400 {object} map[string]string "Некорректные параметры запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs/{id}/lyrics [get]
func (h *Handler) GetSongLyrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		const OP = "SongsHandler.GetSongLyrics"
		log := h.log.With(slog.String("OP", OP))

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("Invalid id query param", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id query param"})
			return
		}

		var filter *models.LyricsInput
		if err := c.ShouldBindQuery(&filter); err != nil {
			log.Error("Invalid query params", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		lyrics, err := h.services.Songs.GetLyrics(ctx, id, filter)
		if err != nil {
			log.Error("Failed to get song lyrics", sl.Err(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get song lyrics"})
			return
		}

		log.Info("Successfully get song lyrics")
		c.JSON(http.StatusOK, gin.H{"songLyrics": lyrics})
	}
}

// DeleteSong Удаление песни
// @Summary Удаление песни
// @Description Удаляет песню по указанному ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 204 "Песня успешно удалена"
// @Failure 400 {object} map[string]string "Некорректный ID"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs/{id} [delete]
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

// UpdateSong Обновление данных песни
// @Summary Обновление данных песни
// @Description Обновляет данные песни по указанному ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.SongUpdateInput true "Данные для обновления"
// @Success 204 "Песня успешно обновлена"
// @Failure 400 {object} map[string]string "Некорректные параметры запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs/{id} [put]
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
			log.Error("Invalid body params", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body parameters"})
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

// AddSong Добавление новой песни
// @Summary Добавление новой песни
// @Description Добавляет новую песню, обогащает данные через внешний API и сохраняет в БД
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.SongInput true "Данные новой песни"
// @Success 200 {object} models.Song "Добавленная песня"
// @Failure 400 {object} map[string]string "Некорректные параметры запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs [post]
func (h *Handler) AddSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		const OP = "SongsHandler.AddSong"
		log := h.log.With(slog.String("OP", OP))

		var inputData *models.SongInput
		if err := c.ShouldBind(&inputData); err != nil {
			log.Error("Invalid body params", sl.Err(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body parameters"})
			return
		}

		songs, err := h.services.Songs.Add(ctx, inputData)
		if err != nil {
			log.Error("Failed to add song", sl.Err(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add song"})
			return
		}

		log.Info("Successfully add song")
		c.JSON(http.StatusOK, songs)
	}
}
