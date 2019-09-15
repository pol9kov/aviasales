package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pol9kov/aviasales/server"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func updateDictionary(t *testing.T, dictionary []string) {
	testServer := httptest.NewServer(http.HandlerFunc(server.LoadHandler))
	defer testServer.Close()

	b, err := json.Marshal(dictionary)
	require.NoError(t, err)

	_, err = http.Post(testServer.URL, "application/json", bytes.NewBuffer(b))
	require.NoError(t, err)
}

func checkAnagram(t *testing.T, word string, expectedAnagrams []string) {
	testServer := httptest.NewServer(http.HandlerFunc(server.GetHandler))
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/get?word=%s", testServer.URL, word))
	require.NoError(t, err)

	var actualAnagrams []string
	json.NewDecoder(resp.Body).Decode(&actualAnagrams)
	resp.Body.Close()
	require.Equal(t, expectedAnagrams, actualAnagrams)
}

func TestAnagram(t *testing.T) {

	updateDictionary(t, []string{"foobar", "aabb", "baba", "boofar", "test"})

	checkAnagram(t,"foobar", []string{"foobar","boofar"})
	checkAnagram(t,"raboof", []string{"foobar","boofar"})
	checkAnagram(t,"abba", []string{"aabb","baba"})
	checkAnagram(t,"test", []string{"test"})
	checkAnagram(t,"qwerty", nil)
}