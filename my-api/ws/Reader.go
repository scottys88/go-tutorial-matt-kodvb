package websocketUtils

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		myString := string(p[:])

		fmt.Printf("Received message: %v: %v", myString, messageType)
	}

}
