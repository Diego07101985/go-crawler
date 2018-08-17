package ormsql

import (
	"go-elasticsearch-example/search-api/api/models"

	"github.com/jinzhu/gorm"
)

var (
	err   error
	db    *gorm.DB
	anime models.AnimeDocument
)

func NewDb(database *gorm.DB) {
	db = database
}

func GetAnimeDocumentById(anime models.AnimeDocument, ID int) {

	db.Find(&anime, ID)
}
