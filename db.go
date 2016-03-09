package main

import (
	"log"

	"encoding/json"
	"os"

	"github.com/jinzhu/gorm"
)

var mail Mail
var bounce Bounce
var delivery Delivery
var notification Notification

type Configuration struct {
	Adapter  string
	User     string
	Database string
	Password string
	Host     string
	Port     string
}

func connectDB() (gorm.DB, error) {

	var configuration Configuration
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&configuration); err != nil {
		log.Println(err)
	}

	db, error := gorm.Open(configuration.Adapter, "user="+configuration.User+" dbname="+configuration.Database+" password="+configuration.Password+" host="+configuration.Host+" port="+configuration.Port+" sslmode=disable")
	return db, error
}

func createTables() {
	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("failed connect DB")
	}

	db.DB()

	db.CreateTable(&mail)
	db.CreateTable(&bounce)
	db.CreateTable(&delivery)
	db.CreateTable(&notification)
}

func deleteTables() {
	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("failed connect DB")
	}

	db.DB()

	db.DropTable(&notification)
	db.DropTable(&mail)
	db.DropTable(&bounce)
	db.DropTable(&delivery)
}

func setForeignKeys() {
	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
		log.Fatal("failed connect DB")
	}

	db.DB()
	db.Model(&notification).AddForeignKey("mail_id", "mails(id)", "RESTRICT", "RESTRICT")
}
