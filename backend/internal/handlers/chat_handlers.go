package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/omnichat-backend/internal/models"
	"github.com/user/omnichat-backend/internal/repository"
)

func GetChats(c *gin.Context) {
	var chats []models.Chat
	if repository.DB != nil {
		repository.DB.Preload("Contact").Order("updated_at desc").Find(&chats)
	}
	c.JSON(http.StatusOK, chats)
}

func GetMessages(c *gin.Context) {
	chatID := c.Param("id")
	var messages []models.Message
	if repository.DB != nil {
		repository.DB.Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages)
	}
	c.JSON(http.StatusOK, messages)
}

type SendMessageRequest struct {
	Text string `json:"text"`
}

// SendMessageFromAgent handles message sending from the UI
func SendMessageFromAgent(c *gin.Context) {
	chatIDStr := c.Param("id")
	chatID, _ := strconv.ParseUint(chatIDStr, 10, 32)
	
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbMsg := models.Message{
		ChatID: uint(chatID),
		Sender: "agent",
		Text:   req.Text,
		Status: "sent",
	}
	
	if repository.DB != nil {
		repository.DB.Create(&dbMsg)
	}

	// In a real app, you'd trigger the external provider (TG/WA) here
	// services.SendMessageToProvider(uint(chatID), req.Text)

	c.JSON(http.StatusOK, dbMsg)
}
