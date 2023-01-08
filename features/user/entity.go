package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
	Name string
	Email string
	HP string
	Address string
	Password string
}

type UserHandler interface{
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactivate() echo.HandlerFunc
} 

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updatedProfile Core) (Core, error)
	Deactivate(token interface{}) (error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(email string) (Core, error)
	Profile(id uint) (Core, error)
	Update(updatedProfile Core) (Core, error)
	Deactivate(id uint) (error)
}