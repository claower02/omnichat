package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/user/omnichat-backend/internal/handlers"
	"github.com/user/omnichat-backend/internal/repository"
)

func main() {
	log.Println(">>> Starting Omnichat Application <<<")
	// Load .env if exists
	godotenv.Load()

	// Initialize DB
	log.Println(">>> Initializing Database...")
	repository.InitDB()
	log.Println(">>> Database Initialized")

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// WebSocket route for frontend
	r.GET("/ws", handlers.ServeWS)

	// API Routes for Frontend
	api := r.Group("/api/v1")
	{
		api.GET("/chats", handlers.GetChats)
		api.GET("/chats/:id/messages", handlers.GetMessages)
		api.POST("/chats/:id/send", handlers.SendMessageFromAgent)
	}

	// Serve static files from frontend build
	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.svg", "./public/favicon.svg")
	r.StaticFile("/icons.svg", "./public/icons.svg")
	
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// API Routes for CRM
	crm := r.Group("/crm/v1")
	{
		crm.POST("/send", handlers.CRMInboundSend) // CRM sends message to chat
		crm.POST("/webhook-config", handlers.SetWebhookConfig) // Setup webhook URL
	}

	// Webhooks from Providers (e.g., Telegram, WhatsApp)
	providers := r.Group("/webhooks")
	{
		providers.POST("/telegram", handlers.TelegramWebhook)
		providers.POST("/whatsapp", handlers.WhatsAppWebhook)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
