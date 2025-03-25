package client

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/gorilla/websocket"
	"github.com/leosalgado/socket/config"
)

func StartClient() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+config.HOST+config.PORT+"/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit.")

	for {
		key, keyCode, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if keyCode == keyboard.KeyEsc {
			fmt.Println("Exiting...")
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(string(key))); err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent key: %q\n", key)
	}
}
