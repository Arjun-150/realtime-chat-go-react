package websocket

import (
	"fmt"

	"github.com/Arjun-150/realtime-chat-go-react/pkg/db" // 🚀 Updated Import
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
			fmt.Println("Clients Added:", len(pool.Clients))

			// 🚀 NEW: Fetch history from MongoDB
			history, err := db.GetHistory()
			if err != nil {
				fmt.Println("Error fetching history:", err)
			} else {
				// 🚀 Send each historical message to ONLY the new client
				for _, msg := range history {
					// We use the local websocket message type here
					client.Conn.WriteJSON(msg)
				}
			}

			pool.broadcast(Message{
				Type:     "system",
				Username: client.Username,
				Body:     "joined",
			})
		case client := <-pool.Unregister:
			if _, ok := pool.Clients[client]; ok {
				delete(pool.Clients, client)
				fmt.Println("Clients Removed:", len(pool.Clients))
				pool.broadcast(Message{
					Type:     "system",
					Username: client.Username,
					Body:     "left",
				})
			}

		case message := <-pool.Broadcast:
			// 🔍 DEBUG 1: See what actually arrived in the pool
			fmt.Printf("DEBUG: Pool received message: %+v\n", message)

			// 🚀 Check for "chat" (Case sensitive! If React sends "Chat", this fails)
			if message.Type == "chat" {
				fmt.Println("DEBUG: Type is 'chat', attempting DB save...")

				// Use the struct exactly as defined in your db package
				err := db.InsertMessage(db.Message{
					Type:     message.Type,
					Username: message.Username,
					Body:     message.Body,
					Time:     message.Time,
				})

				if err != nil {
					fmt.Println("❌ DB Save Error:", err)
				} else {
					fmt.Println("💾 Success: Saved to MongoDB")
				}
			} else {
				fmt.Printf("DEBUG: Message type was '%s', skipping DB save.\n", message.Type)
			}

			// This MUST be here to send messages back to the webpage
			pool.broadcast(message)
		}
	}
}

func (pool *Pool) broadcast(msg Message) {
	for c := range pool.Clients {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("broadcast error:", err)
		}
	}
}
