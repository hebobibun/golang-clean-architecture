package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
	Title string `validate:"required"`
	Year int `validate:"required"`
	Author string `validate:"required"`
	UserID uint
	Owner string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	MyBook() echo.HandlerFunc
	BookList() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	MyBook(token interface{}) ([]Core, error)
	BookList() ([]Core, error)
	Update(token interface{}, bookID uint, updatedData Core) (Core, error)
	Delete(token interface{}, bookID uint) (error)
}

type BookData interface {
	Add(userID uint, newBook Core) (Core, error)
	MyBook(userID uint) ([]Core, error)
	BookList() ([]Core, error)
	Update(userID uint, bookID uint, updatedData Core) (Core, error)
	Delete(userID, bookID uint) (error)
}