package models

import (
	"letsgo/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {

	cfg := config.LoadConfig()

	database, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Sensor{})
	if err != nil {
		panic("Failed to auto-migrate")
	}

	DB = database

	return DB
}
