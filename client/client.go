package client

import (
	"fmt"
	"log"
	"net"

	"github.com/eiannone/keyboard"
)

const (
	HOST = "localhost"
	PORT = ":7740"
	TYPE = "tcp"
)

func StartClient() {
	const (
		HOST = "localhost"
		PORT = ":7740"
		TYPE = "tcp"
	)

	conn, err := net.Dial(TYPE, HOST+PORT)
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
			// break
		}

		_, err = conn.Write([]byte(string(key)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent key: %q\n", key)
	}
}
