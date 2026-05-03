package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/user/omnichat-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println(">>> DATABASE_URL not set. Skipping database connection.")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: Failed to connect to database: %v. App will continue but DB features may fail.", err)
		return
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.Contact{}, &models.Chat{}, &models.Message{}, &models.WebhookConfig{})
	if err != nil {
		log.Printf("ERROR: Failed to migrate database: %v", err)
		return
	}

	fmt.Println("Database connection established and migrated")
}
