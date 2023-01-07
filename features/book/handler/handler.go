package handler

import (
	"go-clean-arch/features/book"
	"go-clean-arch/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookControll struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookControll{
		srv: bs,
	}
}

func (bc *bookControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		cnv := ToCore(input)

		res, err := bc.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("Error insert book :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		finRes := ToResponse(res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Added a new book successfully", finRes))
	}
}

// func (bc *bookControll) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := c.Get("user")
		
// 		input := UpdateBookRequest{}
// 		if err := c.Bind(&input); err != nil {
// 			return c.JSON(http.StatusBadRequest, "Please input correctly")
// 		}

// 		res, err := bc.srv.Update(token)

// 		if err != nil {
// 			return c.JSON(helper.PrintErrorResponse(err.Error()))
// 		}

// 	}
// }
