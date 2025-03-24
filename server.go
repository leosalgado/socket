package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

const(
	HOST = "localhost"
	PORT = ":7740"
	TYPE = "tcp"
)

func main() {

	/* start TCP server */
	listen, err := net.Listen(TYPE, HOST+PORT)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()
	for{
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)	
			os.Exit(1)
		}
		go handleRequest(conn)
	}


	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			fmt.Println("Quitting...")
			break
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
	}

}

func handleRequest(conn net.Conn){

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil{
		log.Fatal(err)
	}
	
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your key is: %v\nReceived time: %v\n", string(buffer[:]), time)
	conn.Write([]byte(responseStr))
	
	conn.Close()
}