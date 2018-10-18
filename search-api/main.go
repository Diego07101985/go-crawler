package main

import (
	"go-crawler/search-api/api/controllers"
	"go-crawler/search-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

var (
	bulk *elastic.BulkService
	err  error
)

func main() {
	r := gin.Default()
	configRoutes(r)
	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func configRoutes(r *gin.Engine) {
	r.Use(middleware.Site())
	r.POST("/animes/elastic", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	r.PUT("/animes/:id", controllers.UpdateAnimeEndpoint)
	r.DELETE("/animes/:id", controllers.DeleteAnimeEndpoint)
	r.GET("/animes/:id", controllers.FindAnimeEndPoint)
	r.GET("/crawler", controllers.ExecuteCrawlerEndpoint)
	r.POST("/animes", controllers.CreateAnimeEndPoint)
}
