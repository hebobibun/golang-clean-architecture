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

func (bd *bookData) BookList() ([]book.Core, error) {
	res := []AllBooks{}
	if err := bd.db.Table("books").Joins("JOIN users ON users.id = books.user_id").Select("books.id, books.title, books.year, books.author, users.name AS owner").Find(&res).Error; err != nil {
		log.Println("Get all books query error : ", err.Error())
		return []book.Core{}, err
	}

	return AllListToCore(res), nil
}

func (bd *bookData) MyBook(userID uint) ([]book.Core, error) {
	res := []Books{}
	if err := bd.db.Where("user_id = ?", userID).Find(&res).Error; err != nil {
		log.Println("Get book by ID query error : ", err.Error())
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

func (bd *bookData) Delete(userID uint, bookID uint) (error) {
	getID := Books{}
	err := bd.db.Where("id = ?", bookID).First(&getID).Error

	if err != nil {
		log.Println("Get book error : ", err.Error())
		return errors.New("Failed to get book data")
	}

	if getID.UserID != userID {
		log.Println("Unauthorized request")
		return errors.New("Unauthorized request")
	}

	qryDelete := bd.db.Delete(&Books{}, bookID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("Failed to delete, data not found")
	}

	return nil
}