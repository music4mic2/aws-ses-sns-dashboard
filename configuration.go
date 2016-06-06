package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Database  Database
	BasicAuth BasicAuth
}

type Database struct {
	Adapter  string
	User     string
	Database string
	Password string
	Host     string
	Port     string
}

type BasicAuth struct {
	User     string
	Password string
}

func ReadConfiguration() Configuration {
	var configuration Configuration
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&configuration); err != nil {
		log.Fatal(err)
	}

	return configuration
}
