package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
	Title string `validate:"required"`
	Year int `validate:"required"`
	Author string `validate:"required"`
	UserID uint
}

type BookHandler interface {
	Add() echo.HandlerFunc
	// Update() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	// Update(token interface{}, bookID int, updatedData Core) (Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	// Update(bookID int, updatedData Core) (Core, error)
}