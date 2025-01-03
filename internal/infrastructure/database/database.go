package database

import (
	"log"

	"github.com/agilistikmal/foodbar-api/internal/foodbar/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.dsn")))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Product{}, &model.HalalData{}, &model.Auth{})

	return db
}
