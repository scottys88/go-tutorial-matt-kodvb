package websocketUtils

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func Writer(conn *websocket.Conn, message []byte) {

	myString := string(message[:])
	messageResponse := fmt.Sprint("websocket.Writer:" + myString)

	if err := conn.WriteMessage(1, []byte(messageResponse)); err != nil {
		log.Println(err)
		return
	}

}
