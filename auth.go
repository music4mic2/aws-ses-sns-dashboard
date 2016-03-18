package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(w http.ResponseWriter, r *http.Request) bool {

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	var configuration Configuration

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&configuration); err != nil {
		log.Fatal(err)
	}

	var basicAuth BasicAuth = configuration.BasicAuth

	return pair[0] == basicAuth.User && pair[1] == basicAuth.Password
}
