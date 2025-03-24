package client

import (
	"fmt"
	"log"
	"net"

	"github.com/eiannone/keyboard"
	"github.com/leosalgado/socket/config"
)

func StartClient() {

	conn, err := net.Dial(config.TYPE, config.HOST+config.PORT)
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

		_, err = conn.Write([]byte(string(key)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent key: %q\n", key)
	}
}
