package handler

import (
	"fmt"
	"go-clean-arch/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll {
		srv: srv,
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusCreated, "Registered a new account Successfully", res))
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		msg := fmt.Sprintf("Login successful. You are now logged in as %s.", res.Name)

		return c.JSON(PrintSuccessReponse(http.StatusOK, msg, res, token))
	}
}

func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)

		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "Displayed your profile successfully", res))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		res, err := uc.srv.Update(token, *ToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, ToResponse(res))
	}
}

func (uc *userControll) Deactivate() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := uc.srv.Deactivate(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "Deactivate your account successfully")
	}
}