package models

import "time"

// Song - структура для хранения песни в БД
type Song struct {
	ID          int    `db:"id" json:"id"`
	Group       string `db:"group" json:"group"`
	Song        string `db:"song" json:"song"`
	ReleaseDate string `db:"release_date" json:"releaseDate"`
	Text        string `db:"text" json:"text"`
	Link        string `db:"link" json:"link"`
}

func (s *Song) GetReleaseDate() string {
	return s.ReleaseDate
}

// SongFilter - структура для фильтрации данных при GET запросах на получение
type SongFilter struct {
	Group       string `form:"group" json:"group"`
	Song        string `form:"song" json:"song"`
	ReleaseDate string `form:"releaseDate" json:"release_date"`
	Text        string `form:"text" json:"text"`
	Link        string `form:"link" json:"link"`
	Offset      int    `form:"offset,default=0" binding:"min=0" json:"offset"`
	Limit       int    `form:"limit,default=0" binding:"min=0" json:"limit"`
}

func (s *SongFilter) GetReleaseDate() string {
	return s.ReleaseDate
}

// SongUpdateInput - структура для обновления данных при PUT запросах на обновление данных
type SongUpdateInput struct {
	Group       string `db:"group" json:"group"`
	Song        string `db:"song" json:"song"`
	ReleaseDate string `db:"release_date" json:"releaseDate"`
	Text        string `db:"text" json:"text"`
	Link        string `db:"link" json:"link"`
}

func (s *SongUpdateInput) GetReleaseDate() string {
	return s.ReleaseDate
}

// SongInput - структура для добавления новой песни при POST запросе
type SongInput struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

// LyricsInput - структура для пагинации куплетов песен, для GET запроса на получение теста песни
type LyricsInput struct {
	Offset int `form:"offset,default=0" binding:"min=0" json:"offset"`
	Limit  int `form:"limit,default=0" binding:"min=0" json:"limit"`
}

type dateFormater interface {
	GetReleaseDate() string
}

func ValidateSongDate(data dateFormater) bool {
	dateStr := data.GetReleaseDate()
	var err error
	if dateStr != "" {
		_, err = time.Parse("2006-01-02", dateStr)
	}
	return err == nil
}
