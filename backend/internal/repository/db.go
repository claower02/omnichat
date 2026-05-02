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
		log.Println("DATABASE_URL not set, using sqlite for local fallback if needed, but assuming postgres")
		dsn = "host=localhost user=postgres password=postgres dbname=omnichat port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.Contact{}, &models.Chat{}, &models.Message{}, &models.WebhookConfig{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connection established and migrated")
}
