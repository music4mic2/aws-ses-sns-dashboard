package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

var mail Mail
var bounce Bounce
var delivery Delivery
var notification Notification
var db *gorm.DB

func dbInstance() *gorm.DB {
	db := connectDB()
	db.DB().SetMaxIdleConns(0)
	db.LogMode(true)

	return db
}

func connectDB() *gorm.DB {
	configuration := ReadConfiguration()

	var database Database = configuration.Database
	db, err := gorm.Open(database.Adapter, "user="+database.User+" dbname="+database.Database+" password="+database.Password+" host="+database.Host+" port="+database.Port+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initialize() {
	db := dbInstance()
	createTables(db)
	setForeignKeys(db)
	setIndexes(db)
}

func createTables(db *gorm.DB) {
	db.CreateTable(&mail)
	db.CreateTable(&bounce)
	db.CreateTable(&delivery)
	db.CreateTable(&notification)
}

func deleteTables(db *gorm.DB) {
	db.DropTable(&notification)
	db.DropTable(&mail)
	db.DropTable(&bounce)
	db.DropTable(&delivery)
}

func setForeignKeys(db *gorm.DB) {
	db.Model(&notification).AddForeignKey("mail_id", "mails(id)", "RESTRICT", "RESTRICT")
}

func setIndexes(db *gorm.DB) {
	db.Model(&notification).AddIndex("index_notification_type", "notification_type")
	db.Model(&mail).AddIndex("index_mail_destination", "destination")
	db.Model(&mail).AddIndex("index_mail_source", "source")
}
