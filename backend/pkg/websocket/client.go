package websocket

import (
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

		msg := Message{
			Type:     "chat",
			Body:     string(p),
			Username: c.Username,
		}

		c.Pool.Broadcast <- msg
		fmt.Println("Received:", msg.Body)
	}
}
