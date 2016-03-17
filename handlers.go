package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

func Notifications(res http.ResponseWriter, req *http.Request) {

	var notification Notification

	if !check(res, req) {
		res.Header().Set("WWW-Authenticate", `Basic realm="beetrack.com"`)
		res.WriteHeader(401)
		res.Write([]byte("401 Unauthorized\n"))
		return
	}

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

	const limit = 100

	user, pass, _ := req.BasicAuth()
	log.Println(user)
	log.Println(pass)

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, Accept")

	if !check(res, req) {
		res.Header().Set("WWW-Authenticate", `Basic realm="beetrack.com"`)
		res.WriteHeader(401)
		res.Write([]byte("401 Unauthorized\n"))
		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	email := req.URL.Query().Get("email")

	if page == 0 {
		page++
	}

	var notifications []Notification

	db := connectDB()
	db.DB()
	db.LogMode(true)

	db.Offset((page - 1) * limit).Limit(limit).Order("created_at asc").Preload("Mail").Preload("Bounce").Find(&notifications)
	if email != "" {
		db.Where("mails.destination LIKE ?", "%"+email+"%").Offset((page - 1) * limit).Limit(limit).Order("created_at asc").Joins("JOIN mails on mails.id = notifications.mail_id").Preload("Mail").Preload("Bounce").Find(&notifications)
	}

	json, err := json.Marshal(notifications)
	if err != nil {
		log.Println(err)
		return
	}
	res.Write(json)
}
