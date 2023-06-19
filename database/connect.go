package database

import (
	"go-jwt/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to database")
	}
	log.Println("Successfully connected to database")
}

func Migration() {
	Instance.AutoMigrate(&model.User{})
	Instance.AutoMigrate(&model.Blog{})
	log.Println("Successfully migrated database")
}
