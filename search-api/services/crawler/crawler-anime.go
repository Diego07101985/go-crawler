package crawler

import (
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
	"go-crawler/search-api/api/repositorys"
	"go-crawler/search-api/api/repositorys/repository-gorm"
	"go-crawler/search-api/infra/db"
	"net/http"
	"strconv"
	"sync"

	"github.com/olivere/elastic"
)

var (
	anime         models.AnimeDocument
	wg            sync.WaitGroup
	slot          = make(chan struct{}, 100)
	channelAnimes = make(chan models.AnimeDocument)
	err           error
	bulkresponse  *elastic.BulkResponse
)

var BASE_URL string = "https://api.jikan.moe/anime/"

func getAnimeInApiExtern(link string, animeChan chan models.AnimeDocument) error {
	r, err := http.Get(link)
	anime := models.AnimeDocument{}

	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&anime)
	animeChan <- anime
	if err != nil {
		return err
	}
	return err
}

func ConsumeAnimes(numberRequest int) {
	db.Init()
	ormsql.NewDb(db.GetDB())

	for i := 0; i < numberRequest; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			createSlotsConcurrencyForAnimes(i, slot)
			if err != nil {
				fmt.Println("Erro")
			}
		}(i)
	}
	for anime := range channelAnimes {
		wg.Add(1)
		go func(animeDocument models.AnimeDocument) {
			if animeDocument.Title != "" {
				ormsql.CreateAnime(animeDocument)
				if _, err := elasticrepo.CreateAnimeDocument(animeDocument); err != nil {
					fmt.Println("Method - ConsumeAnimesErro ao cadastrar anime no Elastic Search ")
				}
			}
		}(anime)
	}
	wg.Wait()
}

func createSlotsConcurrencyForAnimes(numberAnime int, slot chan struct{}) error {
	slot <- struct{}{}
	go func() {
		defer func() { <-slot }()
		fmt.Println(numberAnime)
		err = getAnimeInApiExtern(BASE_URL+strconv.Itoa(numberAnime), channelAnimes)
	}()
	return err
}
