package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	_ "github.com/lib/pq"
)

func Notifications(res http.ResponseWriter, req *http.Request) {
	logRequest(req)
	if !checkAuth(res, req) {
		return
	}
	jsonBody, err := jsonBody(req)
	if err == nil && !isSubscriptionConfirmation(req, jsonBody) {
		var notification Notification

		message := jsonBody["Message"].(string)
		json.Unmarshal([]byte(message), &notification)

		db := dbInstance()
		db.Create(&notification)
	}
}

func NotificationIndex(res http.ResponseWriter, req *http.Request) {

	const limit = 100

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method == "POST" {

		if checkAuth(res, req) {

			page, _ := strconv.Atoi(req.FormValue("page"))
			nType := req.FormValue("type")
			source := req.FormValue("source")
			email := req.FormValue("email")

			if page == 0 {
				page++
			}

			var notifications []Notification

			db := dbInstance()

			chain := db.Offset((page - 1) * limit).Limit(limit).Order("created_at desc").Preload("Mail").Joins("JOIN mails on mails.id = notifications.mail_id")

			if email != "" {
				chain = chain.Where("mails.destination LIKE ?", "%"+email+"%")
			}

			if nType != "" {
				chain = chain.Where("notifications.notification_type = ?", nType)
			}

			if source != "" {
				chain = chain.Where("mails.source LIKE ?", "%"+source+"%")
			}

			chain.Find(&notifications)

			json, err := json.Marshal(notifications)
			if err != nil {
				log.Println(err)
				return
			}
			res.Write(json)
		}
	}
}

func isSubscriptionConfirmation(req *http.Request, jsonBody map[string]interface{}) bool {
	if req.Header.Get("x-amz-sns-message-type") == "" {
		return false
	}

	result := false
	switch req.Header.Get("x-amz-sns-message-type") {
	case "SubscriptionConfirmation":
		log.Println(jsonBody["SubscribeURL"])
		visitURL(jsonBody["SubscribeURL"].(string))
		result = true
	case "Notification":
		log.Println(jsonBody["UnsubscribeURL"])
	}
	return result
}

func visitURL(url string) {
	go func() {
		http.Get(url)
	}()
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

func logRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return
	}

	log.Printf("%s", dump)
}

func jsonBody(req *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(req.Body)

	mapper := make(map[string]interface{})
	err := decoder.Decode(&mapper)

	if err != nil {
		log.Println(err)
	}
	return mapper, err
}
