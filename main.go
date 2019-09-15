package main

import (
	"github.com/pol9kov/aviasales/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/load", server.LoadHandler)
	http.HandleFunc("/get", server.GetHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
