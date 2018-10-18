package models

import (
	"container/list"
	"fmt"
	"go-crawler/search-api/api/repositorys/repository-gorm"
	"time"
)

// AnimeDocument is a representation of a anime

type AnimeDocumentsList struct {
	List *list.List
}

func (e *AnimeDocumentsList) AddToList(anime *repository.AnimeDocument) {
	e.List.PushBack(anime)
}

func (e *AnimeDocumentsList) PrintAllList() {
	for c := e.List.Front(); c != nil; c = c.Next() {
		fmt.Print(c.Value.(*repository.AnimeDocument).ID)
		fmt.Print(c.Value.(*repository.AnimeDocument).Title)
	}
}

// AnimeDocumentRequest is a representation of a anime request
type AnimeDocumentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//AnimeDocumentResponse is a representation of a anime response
type AnimeDocumentResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

//AnimeSearchResponse is a representation of a anime search response
type AnimeSearchResponse struct {
	Time      string                  `json:"time"`
	Hits      string                  `json:"hits"`
	Documents []AnimeDocumentResponse `json:"documents"`
}
