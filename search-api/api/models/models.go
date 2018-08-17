package models

import (
	"time"
)

type PageSearch struct {
	Skip int
	Take int
}

// AnimeDocument is a representation of a anime
type AnimeDocument struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Slug          string
	CreatedAt     time.Time `json:"created_at"`
	Updated       string    `json:"updated_at"`
	Content       string    `json:"content"`
	TitleEnglish  string    `json:"title_english"`
	TitleJapanese string    `json:"title_japanese"`
	Aired         string    `json:"aired"`
	Airing        string    `json:"airing"`
	Background    string    `json:"background"`
	Broadcast     string    `json:"broadcast"`
	Duration      string    `json:"duration"`
	Rank          string    `json:"rank"`
	Episodes      string    `json:"episodes"`
	Favorites     string    `json:"favorites"`
	Image         string    `json:"image_url"`
	Members       string    `json:"members"`
	Popularity    string    `json:"popularity"`
	Rating        string    `json:"rating"`
	Score         string    `json:"score"`
	Source        string    `json:"source"`
	Status        string    `json:"status"`
	ScoredBy      string    `json:"scored_by"`
	Synopsis      string    `json:"synopsis"`
	Type          string    `json:"type"`
	Openings      string    `json:"opening_theme"`
	Endings       string    `json:"ending_theme"`
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
