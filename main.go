package main

import (
	"go-clean-arch/config"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)
}