package routes

import (
	"go-crawler/search-api/api/controllers"
	"go-crawler/search-api/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Site())
	r.POST("/animes/elastic", controllers.CreateDocumentsEndpoint)
	r.GET("/search", controllers.SearchEndpoint)
	r.PUT("/animes/:id", controllers.UpdateAnimeEndpoint)
	r.DELETE("/animes/:id", controllers.DeleteAnimeEndpoint)
	r.GET("/animes/:id", controllers.FindAnimeEndPoint)
	r.GET("/crawler", controllers.ExecuteCrawlerEndpoint)
	r.POST("/animes", controllers.CreateAnimeEndPoint)

	return r
}
