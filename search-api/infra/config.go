package infra

import (
	"go-elasticsearch-example/search-api/api/repositorys"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic"
)

var (
	elasticClient *elastic.Client
	bulk          *elastic.BulkService
	err           error
	db            *gorm.DB
)

func ConfigInitElasticSearchClient() (error, *elastic.Client) {
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	return err, elasticClient
}

func ConfigDBOrmGorm() *gorm.DB {
	db, err = gorm.Open("mysql", "root:33838449@tcp(127.0.0.1:3306)/localhost?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&repositorys.AnimeDocument{})
	return db
}
