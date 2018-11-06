package main

import (
	"go-crawler/search-api/api"
	"log"

	"github.com/olivere/elastic"
)

var (
	bulk *elastic.BulkService
	err  error
)

func main() {
	r := routes.ConfigRoutes()
	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
