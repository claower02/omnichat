package services

import (
	"log"

	"github.com/user/omnichat-backend/internal/handlers"
	"github.com/user/omnichat-backend/internal/models"
	"github.com/user/omnichat-backend/internal/repository"
)

// HandleIncomingMessage processes messages from any source (TG, WA, etc.)
func HandleIncomingMessage(msg IncomingMessage) {
	if repository.DB == nil {
		log.Println("DB not initialized, skipping storage")
		return
	}

	// 1. Find or create Contact
	var contact models.Contact
	repository.DB.Where("phone = ? AND provider = ?", msg.ExternalChatID, msg.Provider).FirstOrCreate(&contact, models.Contact{
		Name:     msg.SenderName,
		Phone:    msg.ExternalChatID,
		Provider: msg.Provider,
	})

	// 2. Find or create Chat
	var chat models.Chat
	repository.DB.Where("contact_id = ?", contact.ID).FirstOrCreate(&chat, models.Chat{
		ContactID:  contact.ID,
		ProviderID: msg.ExternalChatID,
		Status:     "open",
	})

	// 3. Save Message
	dbMsg := models.Message{
		ChatID: chat.ID,
		Sender: "client",
		Text:   msg.Text,
		Status: "delivered",
	}
	repository.DB.Create(&dbMsg)

	// 4. Broadcast via WebSocket to Frontend
	handlers.BroadcastEvent("message.received", dbMsg)
	log.Printf("Message received from %s (%s): %s", msg.SenderName, msg.Provider, msg.Text)
}
