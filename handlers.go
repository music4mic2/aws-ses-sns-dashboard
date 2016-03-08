package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func Notifications(res http.ResponseWriter, req *http.Request) {

	var notification Notification

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&notification); err != nil {
		log.Fatal(err)
	}

	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("failed connect DB")
	}

	db.DB()
	db.LogMode(true)

	log.Println(notification)
	db.Create(&notification)
}
