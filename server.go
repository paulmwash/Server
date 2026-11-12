package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

// open or create sql database
func OpenDB() (*sql.DB, error) {
	db, er := sql.Open("sqlite3", "./chat.db")
	if er != nil {
		return nil, er
	}

	// check if database is accessible
	if er := db.Ping(); er != nil {
		db.Close()
		return nil, er
	}

	log.Println("SQLite database connected successfully")

	return db, nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //r.Header.Get("Origin") == "http://localhost:8090"
	},
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed %v", err)
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
