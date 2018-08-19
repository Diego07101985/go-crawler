package ormsql

import (
	"go-crawler/search-api/api/models"

	"github.com/jinzhu/gorm"
)

var (
	err           error
	db            *gorm.DB
	animeInstance models.AnimeDocument
)

func NewDb(database *gorm.DB) {
	db = database
}

func GetAnimeById(ID string) models.AnimeDocument {
	db.Find(&animeInstance, ID)
	return animeInstance
}

func CreateAnime(anime models.AnimeDocument) models.AnimeDocument {
	if anime = GetAnimeById(anime.ID); anime.ID != "" {
		db.Create(&anime)
	}
	return anime
}

func UpdateAnime(anime models.AnimeDocument, ID string) models.AnimeDocument {

	if animeOld := GetAnimeById(ID); &animeOld != nil {
		return animeOld
	}
	db.Save(&anime)
	return anime
}

func DeleteAnime(anime models.AnimeDocument, ID string) models.AnimeDocument {
	animeDeletar := GetAnimeById(ID)
	if &animeDeletar == nil {
		return animeDeletar
	}

	db.Delete(&animeDeletar)
	return animeDeletar
}
