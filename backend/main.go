package main

import (
	"fmt"
	"net/http"

	"github.com/Arjun-150/realtime-chat-go-react/pkg/db" // 🚀 Import DB
	"github.com/Arjun-150/realtime-chat-go-react/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	// 🚀 NEW: Clear Chat Endpoint
	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Handle CORS
		err := db.ClearAllMessages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Chat Cleared")
	})
}

func main() {
	fmt.Println("Chat App v1.0")

	// 🚀 Initialize MongoDB
	db.InitDB()

	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
