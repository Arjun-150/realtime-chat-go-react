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
			fmt.Println("Received:", message.Body)
			pool.broadcast(message)
		}
	}
}

// helper function
func (pool *Pool) broadcast(msg Message) {
	for c := range pool.Clients {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("broadcast error:", err)
		}
	}
}
