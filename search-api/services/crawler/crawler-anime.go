package crawler

import (
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
	"go-crawler/search-api/api/repositorys/repository-gorm"
	"go-crawler/search-api/infra/db"
	"net/http"
	"strconv"
	"sync"
)

var (
	anime models.AnimeDocument
	wg    sync.WaitGroup
)

var BASE_URL string = "https://api.jikan.moe/anime/"

func GetAnimeInApiExtern(link string, animeChan chan models.AnimeDocument) error {
	r, err := http.Get(link)

	anime := models.AnimeDocument{}

	if err != nil {
		fmt.Println(link, "might be down")
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

func CheckAnime() {

	db.Init()
	ormsql.NewDb(db.GetDB())

	channelAnimes := make(chan models.AnimeDocument)

	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := GetAnimeInApiExtern(BASE_URL+strconv.Itoa(i), channelAnimes)
			if err != nil {
				fmt.Println("Erro")
			}
		}(i)
	}

	for anime := range channelAnimes {
		wg.Add(1)
		fmt.Println(anime.Title)
		fmt.Println(anime.TitleJapanese)
		go func(animeDocument models.AnimeDocument) {
			//	if animeDocument.Title != "" {
			ormsql.CreateAnime(animeDocument)
			//	}
		}(anime)
	}
	wg.Wait()
}
