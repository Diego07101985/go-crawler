package main

import (
	"go-crawler/search-api/api/controllers"
	"go-crawler/search-api/middleware"
	"go-crawler/search-api/services/crawler"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

var (
	bulk *elastic.BulkService
	err  error
)

func main() {
	/*	r := gin.Default()
		configRoutes(r)
		if err = r.Run(":8080"); err != nil {
			log.Fatal(err)
		}*/
	crawler.CheckAnime()
}

func configRoutes(r *gin.Engine) {
	r.Use(middleware.Site())
	r.POST("/document", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	r.GET("/findAnimes/:id", controllers.FindAnimeEndPoint)
	r.POST("/createAnime", controllers.CreateAnimeEndPoint)
}
