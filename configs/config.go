package configs

import (
	"github.com/emoncse/airflow/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB

func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort

	// Connect to the database
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errorDb := Database.AutoMigrate(&models.Payment{})
	if errorDb != nil {
		log.Fatal("Auto migrations failed.:", errorDb)
	}
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}
