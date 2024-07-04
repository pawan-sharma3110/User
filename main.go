package main

import (
	"net/http"
	"user/handler"
)

func main() {
	http.HandleFunc("/registr", handler.RegisterUser)
	http.ListenAndServe(":8080", nil)
}
