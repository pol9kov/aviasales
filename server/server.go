package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pol9kov/aviasales/dictionary"
)

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	var words = []string{}
	if err := json.NewDecoder(r.Body).Decode(&words); err != nil {
		log.Print(err)
		return
	}
	r.Body.Close()

	dictionary.Load(words)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-type", "applciation/json")

	words, ok := r.URL.Query()["word"]
	if !ok || len(words[0]) < 1 {
		log.Println("Url Param 'word' is missing")
		return
	}

	json.NewEncoder(w).Encode(dictionary.Get(words[0]))
}
