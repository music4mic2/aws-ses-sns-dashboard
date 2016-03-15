package main

import (
	"encoding/json"
	"html/template"
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

func NotificationIndex(res http.ResponseWriter, req *http.Request) {
	if !check(res, req) {
		res.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		res.WriteHeader(401)
		res.Write([]byte("401 Unauthorized\n"))
		return
	}

	var notifications []Notification

	db := connectDB()
	db.DB()
	db.LogMode(true)
	db.Preload("Mail").Find(&notifications)

	json, err := json.Marshal(notifications)
	if err != nil {
		log.Println(err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*");
	res.Write(json)
}

func Stylesheets() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
}

//html views
func Dashboard(res http.ResponseWriter, req *http.Request) {
	var notification Notification
	db := connectDB()
	db.DB()

	tmpl, _ := template.ParseFiles("views/dashboard.html")
	tmpl.Execute(res, notification)

	res.Header().Set("Content-Type", "text/html")
}
