package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type User struct {
	Name string
	Age  int
}

func (c *User) Sweet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s, you are %d years old.", c.Name, c.Age)
}

func main() {
	log.Printf("Server starting on ....")
	user := &User{
		Name: "Paul",
		Age:  21,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", user.Sweet)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
