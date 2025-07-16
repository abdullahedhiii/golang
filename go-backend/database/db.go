package database

import (
	"fmt"
	"log"

	"github.com/abdullahedhiii/go-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=abdullah password=edhi dbname=BookStore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})

	fmt.Println("Database connected!")
}
