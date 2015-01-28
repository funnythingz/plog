package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	dbmap.DbConnect()
	DBMigrate()
}

func DBMigrate() {
	log.Println(dbmap.Dbmap.DropTableIfExists(&model.Entry{}))
	log.Println(dbmap.Dbmap.CreateTable(&model.Entry{}))
	log.Println(dbmap.Dbmap.AutoMigrate(&model.Entry{}))
}
