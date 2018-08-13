package models

import "time"

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
