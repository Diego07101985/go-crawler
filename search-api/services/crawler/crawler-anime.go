package crawler

import (
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
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

	if err != nil {
		fmt.Println("Não foi possivel criar o objeto")

	}
	animeChan <- anime
	return err
}

func CheckAnime() {
	channelAnimes := make(chan models.AnimeDocument, 20)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := GetAnimeInApiExtern(BASE_URL+strconv.Itoa(i), channelAnimes)
			if err != nil {
				fmt.Println("Erro")
			}
		}()
	}

	wg.Add(1)
	go func() {
		for i := range channelAnimes {
			///Criar Persistencia no elastic Search e no Mysql
		}
	}()

	wg.Wait()
}
