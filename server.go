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