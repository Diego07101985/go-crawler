package controllers

import (
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
	"go-crawler/search-api/api/repositorys"
	"go-crawler/search-api/api/repositorys/repository-gorm"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

var (
	page *models.PageSearch
	err  error
)

func CreateDocumentsEndpoint(c *gin.Context) {
	var doc models.AnimeDocumentRequest
	if err := c.BindJSON(&doc); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}
	newDocumentAnime := models.AnimeDocument{
		ID:        shortid.MustGenerate(),
		Title:     doc.Title,
		CreatedAt: time.Now().UTC(),
		Content:   doc.Content,
	}
	bulk := elasticrepo.CreateAnimeDocument(newDocumentAnime)
	if _, err := elasticrepo.Execute(c, bulk); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}
	c.Status(http.StatusOK)
}

func FindAnimeEndPoint(c *gin.Context) {
	id := c.Param("id")
	anime := ormsql.GetAnimeById(id)
	c.JSON(200, anime)
}

func SearchEndpoint(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	page := models.PageSearch{Skip: 0, Take: 10}

	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		page.Skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		page.Skip = i
	}
	result, err := elasticrepo.SearchDocument(query, &page)

	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := models.AnimeSearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}
	docs := make([]models.AnimeDocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc models.AnimeDocumentResponse
		json.Unmarshal(*hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
