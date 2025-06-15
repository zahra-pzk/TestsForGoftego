package ws

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin : func (r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Http cannot upgrade to the WS: ",err)
		return
	}
	defer conn.Close()
	for {
		_, userMessage, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Cannot read messages, error: ",err)
			break
		}
		fmt.Println("Received ", string(userMessage))

		systemResponse := fmt.Sprintf("You:\n   '%s' \nSystem:\n   Your message is '%s'", userMessage, userMessage)
		err = conn.WriteMessage(websocket.TextMessage, []byte(systemResponse))
		if err != nil {
			fmt.Println("Cannot write message, error: ",err)
			break
		}
	}
}
