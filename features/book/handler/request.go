package handler

import "go-clean-arch/features/book"

type AddBookRequest struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Author string `json:"author"`
}

type UpdateBookRequest struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Author string `json:"author"`
}

func ToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddBookRequest:
		cnv := data.(AddBookRequest)
		res.Title = cnv.Title
		res.Year = cnv.Year
		res.Author = cnv.Author
	default:
		return nil
	}

	return &res
}
