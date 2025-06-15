package main

import (
	"fmt"
	"net/http"

	"github.com/zahra-pzk/TestsForGoftego/realTimeTest/ws"
)

func main(){
	http.HandleFunc("/ws", ws.Handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}