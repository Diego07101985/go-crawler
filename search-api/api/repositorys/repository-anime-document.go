package repositorys

import (
	"go-elasticsearch-example/search-api/api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

// AnimeDocument is a representation of a anime
type AnimeDocument struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
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
)

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

/*
func (d AnimeDocument) UpdateAnimeDocument() AnimeDocument {

}
*/
/*
func (d AnimeDocument) RemoveAnimeDocument() {

}*/

// InitElastic is init class Cliente in elastic
func InitElastic(erro error, elastic *elastic.Client) {
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
