package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	db.DbLoad()
	migrate()
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&model.Entry{}, &model.Comment{}))
}
