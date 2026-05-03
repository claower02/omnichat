package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/omnichat-backend/internal/services"
)

type TelegramUpdate struct {
	Message struct {
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		From struct {
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Text string `json:"text"`
	} `json:"message"`
}

func TelegramWebhook(c *gin.Context) {
	var update TelegramUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ignored"})
		return
	}

	if update.Message.Text != "" {
		services.HandleIncomingMessage(services.IncomingMessage{
			ExternalChatID: fmt.Sprintf("%d", update.Message.Chat.ID),
			SenderName:     update.Message.From.FirstName,
			Text:           update.Message.Text,
			Provider:       "telegram",
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func WhatsAppWebhook(c *gin.Context) {
	// Logic for WhatsApp (e.g. Meta Webhooks)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
