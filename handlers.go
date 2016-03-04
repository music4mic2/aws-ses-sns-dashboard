package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Bounces(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var bounceType BounceType
	err := decoder.Decode(&bounceType)
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("failed connect DB")
	}

	db.DB()
	log.Println(db.DB().Ping())
	db.CreateTable(&bounceType)
	db.Create(&bounceType)
}

func Complaints(w http.ResponseWriter, r *http.Request) {
}

func Deliveries(res http.ResponseWriter, req *http.Request) {
}

func connectDB() (gorm.DB, error) {
	db, error := gorm.Open("postgres", "user=postgres dbname=notifications password=password host=localhost port=5433 sslmode=disable")
	log.Println(db.DB().Ping())
	return db, error
}
