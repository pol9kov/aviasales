package dictionary

import (
	"github.com/pol9kov/aviasales/sortrunes"
	"log"
	"sync"
)

var (
	dictionary        sync.Map
	wordsToDictionary = make(chan []string)
)

func init() {
	go func() {
		for {
			select {
			case words := <-wordsToDictionary:
				writeToDictionary(words)
			}
		}
	}()
}

func Load(words []string) {
	wordsToDictionary <- words
}

func Get(word string) interface{} {
	v, b := dictionary.Load(leadToStandard(word))

	if b {
		return v
	}
	return nil
}

func writeToDictionary(words []string) {
	for _, w := range words {
		standardizedWord := leadToStandard(w)
		if actual, ok := dictionary.LoadOrStore(standardizedWord, []string{w}); ok {
			v := addToValues(w, actual)
			dictionary.Store(standardizedWord, v)
		}
	}
}

func addToValues(word string, wordsInterface interface{}) interface{} {
	if words, ok := wordsInterface.([]string); ok {
		return append(words, word)
	}
	log.Printf("cann't cast wordsInterface to '[]string'; actual type is '%T'", wordsInterface)
	return []string{word}
}

func leadToStandard(word string) string {
	return sortrunes.SortString(word)
}
