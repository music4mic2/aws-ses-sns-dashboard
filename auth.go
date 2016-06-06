package main

import (
	"encoding/base64"
	"net/http"
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

	configuration := ReadConfiguration()

	basicAuth := configuration.BasicAuth

	return pair[0] == basicAuth.User && pair[1] == basicAuth.Password
}
