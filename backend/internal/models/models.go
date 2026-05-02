package models

import (
	"time"
	"gorm.io/gorm"
)

type Contact struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone" gorm:"index"`
	Email     string         `json:"email"`
	Notes     string         `json:"notes"`
	Provider  string         `json:"provider"` // e.g. "telegram", "whatsapp"
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Chat struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	ContactID  uint           `json:"contactId"`
	Contact    Contact        `gorm:"foreignKey:ContactID" json:"contact"`
	ProviderID string         `json:"providerId"` // external chat/user ID
	Status     string         `json:"status"`     // "open", "closed"
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Messages   []Message      `gorm:"foreignKey:ChatID" json:"messages"`
}

type Message struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	ChatID    uint           `json:"chatId"`
	Sender    string         `json:"sender"`    // "agent" or "client"
	Text      string         `json:"text"`
	Status    string         `json:"status"`    // "sent", "delivered", "read"
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type WebhookConfig struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	URL       string    `json:"url"`
	Secret    string    `json:"secret"` // Secret for signing outbound webhooks to CRM
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
