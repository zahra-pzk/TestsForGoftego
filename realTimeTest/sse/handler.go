package sse

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cach-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok{
		http.Error(w,"Unsupported Streaming", http.StatusInternalServerError)
		return
	}
	for i := 1; i<=10; i++ {
		fmt.Fprintf(w,"data: Message number %d at %s \n\n", i, time.Now().Format("15:04:05"))
		flusher.Flush()
		time.Sleep(2*(time.Second))
	} 
}