package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"user-authentication-with-go/pkg/models"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbName)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Purchase{})
	return db
}
