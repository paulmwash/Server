package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type routing struct {
	Message string `json:"message"`
	Source  string `json:"source"`
}

func (s *Server) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	ra := json.NewEncoder(w).Encode(s)
	if ra != nil {
		log.Printf("Encoding problem %v", ra)
	}
}

func (o *routing) Router(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	er := json.NewEncoder(w).Encode(o)
	if er != nil {
		log.Printf("Encoding problem %v", er)
	}
}

func main() {
	procure := &Server{
		Name:    "Server-Ke",
		Version: "1.0.0",
	}

	dest := &routing{
		Message: "Hello from Kenya",
		Source:  "Kenya",
	}

	data, err := json.Marshal(dest)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	pt, er := os.Create("records.json")
	if er != nil {
		log.Fatalf("Error creating file %v", er)
	}
	defer pt.Close()

	rr := os.WriteFile("records.json", data, 0o644)
	if rr != nil {
		log.Fatalf("Error writing to file %v", rr)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", procure.Test)
	mux.HandleFunc("/nest", dest.Router)

	log.Printf("Starting server on port 8080...")

	logg := log.New(os.Stdout, "http: ", log.Ldate)

	ee := http.ListenAndServe(":8080", mux)
	if ee != nil {
		logg.Fatalf("Server cannot be started %v", ee)
	}
}
