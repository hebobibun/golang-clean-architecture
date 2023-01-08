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
	Show() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Show(token interface{}) ([]Core, error)
	Update(token interface{}, bookID uint, updatedData Core) (Core, error)
}

type BookData interface {
	Add(userID uint, newBook Core) (Core, error)
	Show(userID uint) ([]Core, error)
	Update(userID uint, bookID uint, updatedData Core) (Core, error)
}