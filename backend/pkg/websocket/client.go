package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
	Pool     *Pool
}

type Message struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
	Body     string             `json:"body" bson:"body"`
	Username string             `json:"username,omitempty" bson:"username"`
	Time     string             `json:"time,omitempty" bson:"time"`
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

		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			msg = Message{Type: "chat", Body: string(p)}
		}

		if msg.Username != "" {
			c.Username = msg.Username
		}

		c.Pool.Broadcast <- msg
	}
}
