package crawler

import (
	"container/list"
	"encoding/json"
	"fmt"
	"go-crawler/search-api/api/models"
	"net/http"
	"strconv"
)

var (
	anime models.AnimeDocument
)

var BASE_URL string = "https://api.jikan.moe/anime/"

func GetAnimeInApiExtern(link string, anime models.AnimeDocument) error {
	r, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return err
	}
	defer r.Body.Close()

	fmt.Println(link, "Is up!")
	err = json.NewDecoder(r.Body).Decode(anime)

	return err
}

func CheckAnime() {
	l := list.New()

	for i := 0; i < 10; i++ {
		anime := models.AnimeDocument{}
		go GetAnimeInApiExtern(BASE_URL+strconv.Itoa(i), anime)
		l.PushFront(anime)
	}
}
