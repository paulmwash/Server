package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:8090"
	},
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("You're not using a websocket %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error %v", err)
			break
		}
		log.Printf("Received : %s", string(message))
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error can't write the message")
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", webHandler)
	log.Println("Server running now ...")
	ra := http.ListenAndServe(":8090", nil)
	if ra != nil {
		log.Printf("Error caused breakdown %v", ra)
	}
}
