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
	Background    string    `gorm:"type:text" json:"background"`
	Broadcast     string    `json:"broadcast"`
	Duration      string    `json:"duration"`
	Rank          uint64    `json:"rank"`
	Episodes      uint64    `json:"episodes"`
	Favorites     uint64    `json:"favorites"`
	Image         string    `json:"image_url"`
	Members       uint64    `json:"members"`
	Popularity    uint64    `json:"popularity"`
	Rating        uint64    `json:"rating"`
	Score         uint64    `json:"score"`
	Source        string   `json:"source"`
	Status        string    `json:"status"`
	ScoredBy      uint64    `json:"scored_by"`
	Synopsis      string    `gorm:"type:text" json:"synopsis"`
	Type          string    `json:"type"`
	Openings      string    `json:"opening_theme"`
	Endings       string    `json:"ending_theme"`
	Trailer       string    `json:"trailer_url"`
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
	newTimeString, err := time.Parse(time.RFC3339, d.CreatedAt.String())
	d.CreatedAt = newTimeString
	print(err)

	d.Updated = time.Now().String()
	db.Save(&d)
	return err
}

func (d AnimeDocument) DeleteAnime(ID uint64) error {
	var animeDeletar = AnimeDocument{}
	if animeDeletar = d.GetAnimeById(ID); &animeDeletar == nil {
		return errors.New("Não foi possivel encontrar o anime")
	}
	db.Delete(&animeDeletar)
	return err
}

func Count() int {
	db.Table("anime_documents").Count(&count)
	return count
}
