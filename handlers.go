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
		log.Println(err)
	}

	db := connectDB()
	db.DB()
	db.LogMode(true)
	db.Create(&notification)
}
