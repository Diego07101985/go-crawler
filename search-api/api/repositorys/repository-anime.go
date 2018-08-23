package elasticrepo

import (
	"go-crawler/search-api/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

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

// NewElastic is init class Cliente in elastic
func NewElastic(erro error, elastic *elastic.Client) {
	err, elasticClient = erro, elastic
}

// CreateAnimeDocument is a representation of func createAnime
func CreateAnimeDocument(animeDocument models.AnimeDocument) *elastic.BulkService {
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)

	id := strconv.FormatUint(animeDocument.ID, 64)

	bulk.Add(elastic.NewBulkIndexRequest().Id(id).Doc(animeDocument))
	return bulk
}

func Execute(context *gin.Context, b *elastic.BulkService) (*elastic.BulkResponse, error) {
	bulkresponse, err := b.Do(context.Request.Context())
	return bulkresponse, err
}
