package repositorys

import (
	"go-elasticsearch-example/search-api/api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic"
)

// AnimeDocument is a representation of a anime
type AnimeDocument struct {
	ID            string `json:"id"`
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

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

var (
	elasticClient *elastic.Client
	bulk          *elastic.BulkService
	err           error
	context       *gin.Context
	database      *gorm.DB
)

func NewDb(db *gorm.DB) {
	database = db
}
func SetGinContext(c *gin.Context) {
	context = c
}

func SearchDocument(query string, page *models.PageSearch) (*elastic.SearchResult, error) {
	esQuery := elastic.NewMultiMatchQuery(query, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(page.Skip).Size(page.Take).
		Do(context.Request.Context())

	return result, err
}

// NewElastic is init class Cliente in elastic
func NewElastic(erro error, elastic *elastic.Client) {
	err, elasticClient = erro, elastic
}

// CreateAnimeDocument is a representation of func createAnime
func (d AnimeDocument) CreateAnimeDocument() *elastic.BulkService {
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)

	bulk.Add(elastic.NewBulkIndexRequest().Id(d.ID).Doc(d))
	return bulk
}

func Execute(context *gin.Context, b *elastic.BulkService) (*elastic.BulkResponse, error) {
	bulkresponse, err := b.Do(context.Request.Context())
	return bulkresponse, err
}
