package config

import (
	"myproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "postgres://admin:adm_*@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = db

}
