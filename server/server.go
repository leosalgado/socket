package server

import (
	"fmt"
	"log"
	"net"

	"github.com/leosalgado/socket/config"
)

func StartServer() {

	listen, err := net.Listen(config.TYPE, config.HOST+config.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	fmt.Println("Server is listening on", config.HOST+config.PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		keyPressed := string(buffer)

		fmt.Printf("Server received: %q\n", keyPressed)
	}
}
