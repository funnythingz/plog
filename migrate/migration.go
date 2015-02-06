package main

import (
	"../models"
	"log"
)

func main() {
	DbDevelopmentConnect()
	reset()
	create()
	migrate()
}

func reset() {
	log.Println(Dbmap.DropTableIfExists(&model.Entry{}))
}

func create() {
	log.Println(Dbmap.CreateTable(&model.Entry{}))
}

func migrate() {
	log.Println(Dbmap.AutoMigrate(&model.Entry{}))
}
