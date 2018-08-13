package controllers

import (
	"go-elasticsearch-example/search-api/api/models"
	"go-elasticsearch-example/search-api/api/repositorys"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func CreateDocumentsEndpoint(c *gin.Context) {
	var doc models.AnimeDocumentRequest
	if err := c.BindJSON(&doc); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}
	newDocumentAnime := repositorys.AnimeDocument{
		ID:        shortid.MustGenerate(),
		Title:     doc.Title,
		CreatedAt: time.Now().UTC(),
		Content:   doc.Content,
	}

	bulk := newDocumentAnime.CreateAnimeDocument()
	if _, err := repositorys.Execute(c, bulk); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}
	c.Status(http.StatusOK)
}

/*
func searchEndpoint(c *gin.Context) {
	// Parse request
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query, "title", "content").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := models.SearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}
	// Transform search results before returning them
	docs := make([]models.DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc models.DocumentResponse
		json.Unmarshal(*hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}
*/
func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
