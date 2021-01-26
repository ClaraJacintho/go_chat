package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader to make the standard http connection into a websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all connections for the moment
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messagType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(msg))

		if err := conn.WriteMessage(messagType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// listen websocket indefinitely
	reader(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}
