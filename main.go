package main

import (
	"log"
	"net/http"

	"github.com/pol9kov/aviasales/dictionary"
	"github.com/pol9kov/aviasales/server"
)

func main() {
	http.HandleFunc("/load", server.LoadHandler)
	http.HandleFunc("/get", server.GetHandler)

	go dictionary.LaunchWriter()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
