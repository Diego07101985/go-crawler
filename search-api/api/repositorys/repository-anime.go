package elasticrepo

import (
	"fmt"
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

func DeleteAnimeDocument(animeDocument repository.AnimeDocument) (*elastic.DeleteResponse, error) {
	fmt.Println("DeleteAnimeDocument")
	bulk := elasticClient.Delete().
		Index(elasticIndexName).
		Type(elasticTypeName)

	id := strconv.FormatUint(animeDocument.ID, 16)
	fmt.Println(id)
	deleteResponse, err := bulk.Id(id).Do(context.Request.Context())
	return deleteResponse, err
}

func UpdateAnimeDocument(animeDocument repository.AnimeDocument) (*elastic.UpdateResponse, error) {
	fmt.Println("UpdateAnimeDocument")
	bulk := elasticClient.Update().
		Index(elasticIndexName).
		Type(elasticTypeName)

	id := strconv.FormatUint(animeDocument.ID, 16)

	update, err := bulk.Id(id).
		Upsert(animeDocument).
		Do(context.Request.Context())

	return update, err
}

// CreateAnimeDocument is a representation of func createAnime
func CreateAnimeDocument(animeDocument repository.AnimeDocument) (*elastic.BulkResponse, error) {
	fmt.Println("CreateAnimeDocument")
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)

	id := strconv.FormatUint(animeDocument.ID, 16)
	bulkresponse, err := bulk.Add(elastic.NewBulkIndexRequest().Id(id).Doc(animeDocument)).Do(context.Request.Context())
	return bulkresponse, err
}
