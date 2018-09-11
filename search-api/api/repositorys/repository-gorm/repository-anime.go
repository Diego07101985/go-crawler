package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	db    *gorm.DB
	count int
	err   error
)

func NewDb(database *gorm.DB) {
	db = database
}

type AnimeDocument struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Slug          string
	CreatedAt     time.Time `json:"created_at"`
	Updated       string    `json:"updated_at"`
	Content       string    `json:"content"`
	TitleEnglish  string    `json:"title_english"`
	TitleJapanese string    `json:"title_japanese"`
	Aired         string    `json:"aired"`
	Airing        string    `json:"airing"`
	Background    string    `json:"background"`
	Broadcast     string    `json:"broadcast"`
	Duration      string    `json:"duration"`
	Rank          string    `json:"rank"`
	Episodes      string    `json:"episodes"`
	Favorites     string    `json:"favorites"`
	Image         string    `json:"image_url"`
	Members       string    `json:"members"`
	Popularity    string    `json:"popularity"`
	Rating        string    `json:"rating"`
	Score         string    `json:"score"`
	Source        string    `json:"source"`
	Status        string    `json:"status"`
	ScoredBy      string    `json:"scored_by"`
	Synopsis      string    `json:"synopsis"`
	Type          string    `json:"type"`
	Openings      string    `json:"opening_theme"`
	Endings       string    `json:"ending_theme"`
}

func (d AnimeDocument) GetAnimeById(ID uint64) AnimeDocument {
	db.Find(&d, ID)
	return d
}

func (d AnimeDocument) CreateAnime() uint64 {
	if d = d.GetAnimeById(d.ID); &d != nil {
		db.Create(&d)
	}
	return d.ID
}

func (d AnimeDocument) UpdateAnime(ID uint64) error {
	if animeOld := d.GetAnimeById(ID); &animeOld == nil {
		return errors.New("Não foi possivel encontrar o anime")
	}
	db.Save(&d)
	return err
}

func (d AnimeDocument) DeleteAnime(ID uint64) error {
	var animeDeletar = AnimeDocument{}
	if animeDeletar = d.GetAnimeById(ID); &animeDeletar == nil {
		return errors.New("Não existe o anime selecionado")
	}
	db.Delete(animeDeletar)
	return err

}

func Count() int {
	db.Table("anime_documents").Count(&count)
	return count
}
