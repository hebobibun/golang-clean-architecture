package handler

import "go-clean-arch/features/book"

type BookResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Author string`json:"author"`
}

func ToResponse(data book.Core) BookResponse {
	return BookResponse{
		ID: data.ID,
		Title: data.Title,
		Year: data.Year,
		Author: data.Author,
	}
}