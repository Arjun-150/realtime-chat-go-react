package websocket

import (
	"fmt"

	"github.com/Arjun-150/realtime-chat-go-react/pkg/db"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			history, _ := db.GetHistory()
			for _, msg := range history {
				// Convert db.Message to websocket.Message
				client.Conn.WriteJSON(Message{
					ID:       msg.ID,
					Type:     msg.Type,
					Username: msg.Username,
					Body:     msg.Body,
					Time:     msg.Time,
				})
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)

		case message := <-pool.Broadcast:
			// 🚀 HANDLE DELETE
			if message.Type == "delete" {
				db.DeleteMessageByID(message.Body) // Body is the ID string
				pool.broadcast(message)
				continue
			}

			// 🚀 HANDLE CHAT
			if message.Type == "chat" {
				// Save to DB and get the message BACK with its new ID
				savedMsg, err := db.InsertMessage(db.Message{
					Type:     message.Type,
					Username: message.Username,
					Body:     message.Body,
					Time:     message.Time,
				})

				if err == nil {
					// Update our broadcast message with the real ID from Mongo
					message.ID = savedMsg.ID
					fmt.Println("💾 Saved to DB with ID:", message.ID)
				}
			}

			pool.broadcast(message)
		}
	}
}

func (pool *Pool) broadcast(msg Message) {
	for c := range pool.Clients {
		c.Conn.WriteJSON(msg)
	}
}
