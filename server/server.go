package server

import (
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = ":7740"
	TYPE = "tcp"
)

func StartServer() {

	listen, err := net.Listen(TYPE, HOST+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	fmt.Println("Server is listening on", HOST+PORT)

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
