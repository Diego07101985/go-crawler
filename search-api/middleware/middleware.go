package middleware

import (
	"go-elasticsearch-example/search-api/api/repositorys"
	"go-elasticsearch-example/search-api/infra"

	"github.com/gin-gonic/gin"
)

func Site() gin.HandlerFunc {
	return func(c *gin.Context) {
		repositorys.InitElastic(infra.ConfigInitElasticSearchClient())
		repositorys.SetGinContext(c)
		c.Next()
	}
}
