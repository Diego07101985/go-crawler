package middleware

import (
	"go-elasticsearch-example/search-api/api/repositorys"
	"go-elasticsearch-example/search-api/api/repositorys/repository-gorm"
	"go-elasticsearch-example/search-api/infra"
	"go-elasticsearch-example/search-api/infra/db"

	"github.com/gin-gonic/gin"
)

func Site() gin.HandlerFunc {
	return func(c *gin.Context) {
		elasticrepo.NewElastic(infra.ConfigInitElasticSearchClient())
		elasticrepo.SetGinContext(c)
		ormsql.NewDb(db.GetDB())
		c.Next()
	}
}
