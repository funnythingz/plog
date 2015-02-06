package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	DbTestConnect()
	//reset()
	//create()
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
