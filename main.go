package main

import (
	"go-clean-arch/config"
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

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(config.JWTKey)))

	auth.GET("/profile", userHdl.Profile())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}