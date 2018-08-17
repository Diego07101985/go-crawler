package infra

import (
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
