package middleware

import (
	"go-crawler/search-api/api/repositorys"
	"go-crawler/search-api/api/repositorys/repository-gorm"
	"go-crawler/search-api/infra"
	"go-crawler/search-api/infra/db"

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
