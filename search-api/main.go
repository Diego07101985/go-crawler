package main

import (
	"go-elasticsearch-example/search-api/api/controllers"
	"go-elasticsearch-example/search-api/middleware"
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
	r.POST("/documents", controllers.CreateDocumentsEndpoint)
	//	r.GET("/search", searchEndpoint)
}
