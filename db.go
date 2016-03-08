package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

var mail Mail
var bounce Bounce
var delivery Delivery
var notification Notification

func connectDB() (gorm.DB, error) {
	db, error := gorm.Open("postgres", "user=postgres dbname=notifications password=password host=localhost port=5433 sslmode=disable")
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

	db.Model(&Notification{}).AddForeignKey("delivery_id", "deliveries(id)", "RESTRICT", "RESTRICT")
	db.Model(&Notification{}).AddForeignKey("mail_id", "mails(id)", "RESTRICT", "RESTRICT")
}
