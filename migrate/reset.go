package main

import (
	"../db"
	"../models"
	"log"
)

func main() {
	db.DbLoad()
	createDatabase()
	dropTable()
	createTable()
}

func createDatabase() {
	log.Println(db.Dbmap.Exec("create database plog"))
}

func dropTable() {
	log.Println(db.Dbmap.DropTableIfExists(&model.Entry{}))
	log.Println(db.Dbmap.DropTableIfExists(&model.Comment{}))
}

func createTable() {
	log.Println(db.Dbmap.CreateTable(&model.Entry{}))
	log.Println(db.Dbmap.CreateTable(&model.Comment{}))
}
