package handler

import (
	"go-clean-arch/features/book"
	"go-clean-arch/helper"
	"log"
	"net/http"
	"strconv"

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

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Added a new book successfully", ToResponse(res)))
	}
}

func (bc *bookControll) BookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bc.srv.BookList()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Displayed all books successfully", AllListCoreToResp(res)))
	}
}

func (bc *bookControll) MyBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := bc.srv.MyBook(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Displayed your books successfully", ListCoreToResp(res)))
	}
}

func (bc *bookControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		
		paramID := c.Param("id")
		bookID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("Convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please input number only",
			})
		}

		input := UpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		res, err := bc.srv.Update(token, uint(bookID), *ToCore(input))

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Updated a book successfully", res))
	}
}

func (bc *bookControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")
		bookID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("Convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please input number only",
			})
		}

		err = bc.srv.Delete(token, uint(bookID))
		
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "Deleted a book successfully")
	}
}
