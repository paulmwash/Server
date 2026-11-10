package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade  websocket connection %v", err)
		return
	}
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Hello Paul")); err != nil {
		log.Printf("Error occurred %v", err)
		return
	}
	log.Printf("Message sent successful")
}

func main() {
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	http.HandleFunc("/ws", wsHandler)
	log.Printf("Starting server at port 8080")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	<-interrupt
}
