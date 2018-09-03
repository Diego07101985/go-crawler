package elasticrepo

import (
	repository "go-crawler/search-api/api/repositorys/repository-gorm"
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

type PageSearch struct {
	Skip int
	Take int
}

func SetGinContext(c *gin.Context) {
	context = c
}

func SearchDocument(query string, page *PageSearch) (*elastic.SearchResult, error) {
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
func CreateAnimeDocument(animeDocument repository.AnimeDocument) (*elastic.BulkResponse, error) {
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)

	id := strconv.FormatUint(animeDocument.ID, 64)

	bulk.Add(elastic.NewBulkIndexRequest().Id(id).Doc(animeDocument))
	bulkresponse, err := execute(context, bulk)
	return bulkresponse, err
}

func execute(context *gin.Context, b *elastic.BulkService) (*elastic.BulkResponse, error) {
	bulkresponse, err := b.Do(context.Request.Context())
	return bulkresponse, err
}
