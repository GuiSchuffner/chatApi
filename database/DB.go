package database

import (
	"fmt"
	"log"
	"os"

	"github.com/GuiSchuffner/chatApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSllmode)
	DB, err = gorm.Open(postgres.Open(config))
	if err != nil {
		log.Panic("Error opening database")
	}
	DB.Logger = DB.Logger.LogMode(logger.Info)
	DB.AutoMigrate(&models.User{})
}
