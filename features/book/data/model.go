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

type AllBooks struct {
	ID uint
	Title string
	Year int
	Author string
	Owner string
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

func (dataModel *Books) ModelsToCore() book.Core { 
	return book.Core{
		ID: dataModel.ID,
		Title: dataModel.Title,
		Year: dataModel.Year,
		Author: dataModel.Author,
		UserID: dataModel.UserID,
	}
}

func ListToCore(data []Books) []book.Core {
	var dataCore []book.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func (dataModel *AllBooks) AllModelsToCore() book.Core { 
	return book.Core{
		ID: dataModel.ID,
		Title: dataModel.Title,
		Year: dataModel.Year,
		Author: dataModel.Author,
		Owner: dataModel.Owner,
	}
}

func AllListToCore(data []AllBooks) []book.Core {
	var dataCore []book.Core
	for _, v := range data {
		dataCore = append(dataCore, v.AllModelsToCore())
	}
	return dataCore
}