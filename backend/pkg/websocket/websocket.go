package websocket

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

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil

}

func Reader(conn *websocket.Conn) {
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
