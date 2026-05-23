package websocket

import "fmt"

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

			// notify everyone
			for c := range pool.Clients {
				err := c.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
				if err != nil {
					fmt.Println("write error:", err)
				}
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)

			fmt.Println("Clients Removed:", len(pool.Clients))

			for c := range pool.Clients {
				c.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Left...",
				})
			}

		case message := <-pool.Broadcast:
			fmt.Println("Received:", message.Body)

			for c := range pool.Clients {
				err := c.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println("broadcast error:", err)
				}
			}
		}
	}
}
