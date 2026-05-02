package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all for dev
	},
}

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan interface{}
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

var WSHub = &Hub{
	clients:    make(map[*websocket.Conn]bool),
	broadcast:  make(chan interface{}),
	register:   make(chan *websocket.Conn),
	unregister: make(chan *websocket.Conn),
}

func init() {
	go WSHub.run()
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Printf("WS error: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// BroadcastEvent sends a message to all connected React clients
func BroadcastEvent(event string, payload interface{}) {
	WSHub.broadcast <- map[string]interface{}{
		"event":   event,
		"payload": payload,
	}
}

func ServeWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	WSHub.register <- conn

	// Basic read loop to keep connection alive and detect close
	go func() {
		defer func() {
			WSHub.unregister <- conn
		}()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()
}
