package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TelegramWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func WhatsAppWebhook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
