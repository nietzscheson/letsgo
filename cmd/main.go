package main

import (
	"letsgo/config"
	"letsgo/internal/app"
	"letsgo/internal/models"
)

func main() {

	cfg := config.LoadConfig()

	router := app.SetupRouter()
	models.ConnectDatabase()
	router.Run(cfg.ServerAddress)
}
