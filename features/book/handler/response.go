package handler

import "go-clean-arch/features/book"

type BookResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Author string`json:"author"`
	UserID uint`json:"user_id"`
}

func ToResponse(data book.Core) BookResponse {
	return BookResponse{
		ID: data.ID,
		Title: data.Title,
		Year: data.Year,
		Author: data.Author,
		UserID: data.UserID,
	}
}

func CoresToResponse(dataCore book.Core) BookResponse {
	return BookResponse{
		ID: dataCore.ID,
		Title: dataCore.Title,
		Year: dataCore.Year,
		Author: dataCore.Author,
		UserID: dataCore.UserID,
	}
}

func ListCoreToResp(data []book.Core) []BookResponse{
	var dataResp []BookResponse
	for _, v := range data {
		dataResp = append(dataResp, CoresToResponse(v))
	}
	return dataResp
}

