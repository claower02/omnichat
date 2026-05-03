package services

type IncomingMessage struct {
	ExternalChatID string
	SenderName     string
	Text           string
	Provider       string
}

type Provider interface {
	SendMessage(externalChatID string, text string) error
	GetName() string
}
