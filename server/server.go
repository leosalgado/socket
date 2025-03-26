package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/leosalgado/socket/config"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartServer() {
	fmt.Println("Server listening on ws://" + config.HOST + config.PORT + "/ws")

	http.HandleFunc("/ws", handleConnections)
	log.Fatal(http.ListenAndServe(config.PORT, nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Received %s \n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
