package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/omnichat-backend/internal/models"
	"github.com/user/omnichat-backend/internal/repository"
)

func GetChats(c *gin.Context) {
	var chats []models.Chat
	repository.DB.Preload("Contact").Find(&chats)
	c.JSON(http.StatusOK, chats)
}

func GetMessages(c *gin.Context) {
	chatID := c.Param("id")
	var messages []models.Message
	repository.DB.Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages)
	c.JSON(http.StatusOK, messages)
}

type SendMessageRequest struct {
	Text string `json:"text"`
}

// SendMessageFromAgent handles message sending from the UI
func SendMessageFromAgent(c *gin.Context) {
	chatID := c.Param("id")
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For now just mock the response
	c.JSON(http.StatusOK, gin.H{"status": "queued", "chatId": chatID, "text": req.Text})
}
