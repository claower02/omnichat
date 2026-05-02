package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CRMInboundSend receives a message from CRM and sends it to the chat
func CRMInboundSend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// SetWebhookConfig configures where to send updates to the CRM
func SetWebhookConfig(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
