package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
	Pool     *Pool
}

type Message struct {
	Type     string `json:"type"`
	Body     string `json:"body"`
	Username string `json:"username,omitempty"`
	Time     string `json:"time,omitempty"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 🚀 1. Create an empty message struct
		var msg Message

		// 🚀 2. Decode the JSON payload 'p' into our struct
		if err := json.Unmarshal(p, &msg); err != nil {
			fmt.Println("Error decoding JSON:", err)
			// If it's not JSON, treat it as a plain body (fallback)
			msg = Message{
				Type: "chat",
				Body: string(p),
			}
		}

		// 🚀 3. Sync the client's username if it was sent in the message
		if msg.Username != "" {
			c.Username = msg.Username
		}

		c.Pool.Broadcast <- msg
		fmt.Printf("Message Received: %+v\n", msg)
	}
}
