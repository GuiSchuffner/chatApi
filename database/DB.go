package database

import (
	"log"
	"os"

	"github.com/GuiSchuffner/chatApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSllmode := os.Getenv("DB_SSLMODE")
	config := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(config))
	if err != nil {
		log.Panic("Error opening database")
	}
	DB.AutoMigrate(&models.User{})
}
