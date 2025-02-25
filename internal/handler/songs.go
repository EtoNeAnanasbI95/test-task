package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetSongs() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO сделать пагинацию
		//group := c.Query("group")
		//song := c.Query("song")
		//page := c.DefaultQuery("page", "1")
		//limit := c.DefaultQuery("limit", "10")
	}
}

func (h *Handler) GetSongLyrics() gin.HandlerFunc {
	return func(context *gin.Context) {
		panic("implement me")
	}
}

func (h *Handler) DeleteSong() gin.HandlerFunc {
	return func(context *gin.Context) {
		panic("implement me")
	}
}

func (h *Handler) UpdateSong() gin.HandlerFunc {
	return func(context *gin.Context) {
		panic("implement me")
	}
}

func (h *Handler) AddSong() gin.HandlerFunc {
	return func(context *gin.Context) {
		panic("implement me")
	}
}
