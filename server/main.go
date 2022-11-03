package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	fmt.Fprintf(w, "Hello %s", msg)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
}
