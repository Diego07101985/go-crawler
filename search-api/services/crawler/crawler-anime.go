package crawler

import (
	"fmt"
	"net/http"
)

func CheckAnime(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "Is up!")
	c <- link
}
