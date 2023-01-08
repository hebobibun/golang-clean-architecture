package services

import (
	"errors"
	"go-clean-arch/features/book"
	"go-clean-arch/helper"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type bookSrv struct {
	data book.BookData
	vld  *validator.Validate
}

func New(d book.BookData) book.BookService {
	return &bookSrv{
		data: d,
		vld:  validator.New(),
	}
}

func (bs *bookSrv) Add(token interface{}, newBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("User not found")
	}

	err := bs.vld.Struct(newBook)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return book.Core{}, errors.New("Please input correctly")
	}

	res, err := bs.data.Add(uint(userID), newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Book not found"
		} else {
			msg = "There's a server eror"
		}
		return book.Core{}, errors.New(msg)
	}

	res.UserID = uint(userID)

	return res, nil
}

func (bs *bookSrv) BookList() ([]book.Core, error) {
	res, err := bs.data.BookList()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Book not found"
		} else {
			msg = "There's a server eror"
		}
		return []book.Core{}, errors.New(msg)
	}

	return res, nil
}

func (bs *bookSrv) MyBook(token interface{}) ([]book.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return []book.Core{}, errors.New("User not found")
	}

	res, err := bs.data.MyBook(uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Book not found"
		} else {
			msg = "There's a server eror"
		}
		return []book.Core{}, errors.New(msg)
	}

	return res, nil
}

func (bs *bookSrv) Update(token interface{}, bookID uint, updatedData book.Core) (book.Core, error) {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return book.Core{}, errors.New("Data not found")
	}

	res, err := bs.data.Update(uint(id), bookID, updatedData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "Unauthorized") {
			msg = "Unauthorized request"
		} else {
			msg = "There is a problem with the server"
		}
		return book.Core{}, errors.New(msg)
	}

	res.ID = bookID
	res.UserID = uint(id)

	return res, nil
}