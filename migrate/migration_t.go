package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	db.DbTestConnect()
	reset()
	create()
	migrate()
}

func reset() {
	log.Println(db.Dbmap.DropTableIfExists(&model.Entry{}))
}

func create() {
	log.Println(db.Dbmap.CreateTable(&model.Entry{}))
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&model.Entry{}))
}
