package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

func Notifications(res http.ResponseWriter, req *http.Request) {

	if checkAuth(res, req) {
		body, _ := ioutil.ReadAll(req.Body)

		mapper := make(map[string]string)
		e := json.Unmarshal(body, &mapper)
		log.Println(e)

		var notification Notification

		message := mapper["Message"]
		json.Unmarshal([]byte(message), &notification)

		db := connectDB()
		db.DB()
		db.LogMode(true)
		db.Create(&notification)
	}
}

func NotificationIndex(res http.ResponseWriter, req *http.Request) {

	const limit = 100

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	res.Header().Set("Access-Control-Allow-Origin", "http://localhost")

	if req.Method == "GET" {

		if checkAuth(res, req) {

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
	}
}

func SubscriptionConfirmation(res http.ResponseWriter, req *http.Request) bool {

	if req.Header.Get("x-amz-sns-message-type") != "" {
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Println(err)
		}

		mapper := make(map[string]interface{})
		e := json.Unmarshal(body, &mapper)

		log.Println(e)
		log.Println(mapper)

		switch req.Header.Get("x-amz-sns-message-type") {
		case "SubscriptionConfirmation":
			log.Println(mapper["SubscribeURL"])
			return true
		case "Notification":
			log.Println(mapper["UnsubscribeURL"])

		}
		return true
	} else {
		return false
	}
}

func checkAuth(res http.ResponseWriter, req *http.Request) bool {
	if !check(res, req) {
		res.Header().Set("WWW-Authenticate", `Basic realm="beetrack.com"`)
		res.WriteHeader(401)
		res.Write([]byte("401 Unauthorized\n"))
		return false
	} else {
		return true
	}
}
