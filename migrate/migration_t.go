package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	dbmap.DbTestConnect()
	reset()
	create()
	migrate()
}

func reset() {
	log.Println(dbmap.Dbmap.DropTableIfExists(&model.Entry{}))
}

func create() {
	log.Println(dbmap.Dbmap.CreateTable(&model.Entry{}))
}

func migrate() {
	log.Println(dbmap.Dbmap.AutoMigrate(&model.Entry{}))
}
