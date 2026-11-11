package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	checkOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:8090"
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println("Error occurred %v", err)
		return 
	}
	defer conn.Close()

	for {
		messageType,message,err := conn.ReadMessage()
		if err != nil {
			log.Println("Error %v", err)
			break
		}
		log.Println("Received : %s",string(message))
		err = conn.WriteMessage(messageType,message)
		if err != nil {
			log.Println("Error can't write the message")
			break
		}
	}
}

func main() {
	
}