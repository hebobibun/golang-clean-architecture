package data

import (
	"go-clean-arch/features/book"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title string
	Year int
	Author string
	UserID uint
}

func ToCore(data Books) book.Core {
	return book.Core{
		ID: data.ID,
		Title: data.Title,
		Year: data.Year,
		Author: data.Author,
		UserID: data.UserID,
	}
}

func CoreToData(data book.Core) Books {
	return Books{
		Model: gorm.Model{ID: data.ID},
		Title: data.Title,
		Year: data.Year,
		Author: data.Author,
		UserID: data.UserID,
	}
}