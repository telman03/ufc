package db

import (
//	"fmt"
	"log"
	"os"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL") // Use Digital Oceon

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connected successfully!")
}