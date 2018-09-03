package controllers

import (
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
	"go-crawler/search-api/api/repositorys"
	repository "go-crawler/search-api/api/repositorys/repository-gorm"
	"go-crawler/search-api/services/crawler"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

var (
	page             *elasticrepo.PageSearch
	err              error
	newDocumentAnime repository.AnimeDocument
	animeFound       repository.AnimeDocument
	anime            repository.AnimeDocument = repository.AnimeDocument{}
	doc              models.AnimeDocumentRequest
)

func ExecuteCrawlerEndpoint(c *gin.Context) {
	animesPersist := crawler.ConsumeAnimes(100)
	if animesPersist > 0 {
		c.JSON(http.StatusCreated, animesPersist)
	}
}

func CreateDocumentsEndpoint(c *gin.Context) {
	var doc models.AnimeDocumentRequest

	if err := c.BindJSON(&doc); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	if intID, err := strconv.ParseUint(shortid.MustGenerate(), 10, 64); err != nil {
		var doc models.AnimeDocumentRequest

		newDocumentAnime = repository.AnimeDocument{
			ID:        intID,
			Title:     doc.Title,
			CreatedAt: time.Now().UTC(),
			Content:   doc.Content,
		}
	} else {
		errorResponse(c, http.StatusBadRequest, "Id esta no formato incorreto")
	}

	if _, err := elasticrepo.CreateAnimeDocument(newDocumentAnime); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}
	c.Status(http.StatusOK)
}

func FindAnimeEndPoint(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err == nil {
		errorResponse(c, http.StatusBadRequest, "Id esta no formato incorreto")
	}

	if animeFound := anime.GetAnimeById(uid); &animeFound != nil {
		errorResponse(c, http.StatusNoContent, "Anime nao encontrado")
	}
	c.JSON(200, &animeFound)
}

func CreateAnimeEndPoint(c *gin.Context) {
	if err := c.BindJSON(&anime); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	if anime.ID != 0 {
		errorResponse(c, http.StatusConflict, "O anime ja existe")
	}

	id := anime.CreateAnime()
	c.JSON(http.StatusCreated, id)
}

func SearchEndpoint(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	page := elasticrepo.PageSearch{Skip: 0, Take: 10}

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
