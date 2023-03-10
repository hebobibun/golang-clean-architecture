package main

import (
	"go-clean-arch/config"
	bd "go-clean-arch/features/book/data"
	bhl "go-clean-arch/features/book/handler"
	bsrv "go-clean-arch/features/book/services"
	"go-clean-arch/features/user/data"
	"go-clean-arch/features/user/handler"
	"go-clean-arch/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	bookData := bd.New(db)
	bookSrv := bsrv.New(bookData)
	bookHdl := bhl.New(bookSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/books", bookHdl.BookList())

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(config.JWTKey)))

	auth.GET("/profile", userHdl.Profile())
	auth.PATCH("/profile/update", userHdl.Update())
	auth.DELETE("/profile/deactivate", userHdl.Deactivate())

	auth.GET("/books/my", bookHdl.MyBook())
	auth.POST("/books/add", bookHdl.Add())
	auth.PATCH("/books/update/:id", bookHdl.Update())
	auth.DELETE("/books/delete/:id", bookHdl.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}