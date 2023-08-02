package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseDB *gorm.DB

func InitDB() *gorm.DB{
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_CONNECTION")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}


func GetDB() *gorm.DB {
	return DatabaseDB
}

