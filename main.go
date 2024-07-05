package main

import (
	"log"
	"net/http"
	"user/handler"
)

func main() {
	http.HandleFunc("/register", handler.RegisterUser)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
