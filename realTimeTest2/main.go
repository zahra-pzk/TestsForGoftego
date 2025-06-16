package main

import (
	"fmt"
	"net/http"
	"github.com/zahra-pzk/TestsForGoftego/realTimeTest2/ws"
)

func main(){
	http.HandleFunc("/ws", ws.Handler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}