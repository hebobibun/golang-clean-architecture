package data

import (
	"errors"
	"go-clean-arch/features/book"
	"log"

	"gorm.io/gorm"
)


type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID uint, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		return book.Core{}, err
	}

	newBook.ID = cnv.ID

	return newBook, nil
}

func (bd *bookData) Show(userID uint) ([]book.Core, error) {
	res := []Books{}
	if err := bd.db.Where("user_id = ?", userID).Find(&res).Error; err != nil {
		log.Println("Get profile by ID query error : ", err.Error())
		return []book.Core{}, err
	}

	return ListToCore(res), nil
}

func (bd *bookData) Update(userID uint, bookID uint, updatedData book.Core) (book.Core, error) {
	getID := Books{}
	err := bd.db.Where("id = ?", bookID).First(&getID).Error
	
	if err != nil {
		log.Println("Get book error : ", err.Error())
		return book.Core{}, err
	}
	
	if getID.UserID != userID {
		log.Println("Unauthorized request")
		return book.Core{}, errors.New("Unauthorized request")
	}
	
	cnv := CoreToData(updatedData)
	qryUpdate := bd.db.Where("id = ?", bookID).Updates(&cnv)
	if qryUpdate.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return book.Core{}, errors.New("not found")
	}

	if err := qryUpdate.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return book.Core{}, err
	}

	return updatedData, nil
}