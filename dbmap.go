package main

import (
	"./helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "log"
	_ "reflect"
)

var Dbmap gorm.DB

func DbOpen(
	adapter string,
	username string,
	password string,
	database string,
	encoding string) {

	Dbmap, _ = gorm.Open(adapter, username+":"+password+"@/"+database+"?charset="+encoding+"&parseTime=True")

	Dbmap.DB()

	Dbmap.DB().Ping()
	Dbmap.DB().SetMaxIdleConns(10)
	Dbmap.DB().SetMaxOpenConns(100)

	Dbmap.SingularTable(true)
}

func DbConfig() map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	config, _ := ioutil.ReadFile("database.yml")
	err := yaml.Unmarshal([]byte(config), &m)
	helper.Check(err)
	return m
}

func DbDevelopmentConnect() {
	DbOpen(
		DbConfig()["development"].(map[interface{}]interface{})["adapter"].(string),
		DbConfig()["development"].(map[interface{}]interface{})["username"].(string),
		DbConfig()["development"].(map[interface{}]interface{})["password"].(string),
		DbConfig()["development"].(map[interface{}]interface{})["database"].(string),
		DbConfig()["development"].(map[interface{}]interface{})["encoding"].(string),
	)
}

func DbTestConnect() {
	DbOpen(
		DbConfig()["test"].(map[interface{}]interface{})["adapter"].(string),
		DbConfig()["test"].(map[interface{}]interface{})["username"].(string),
		DbConfig()["test"].(map[interface{}]interface{})["password"].(string),
		DbConfig()["test"].(map[interface{}]interface{})["database"].(string),
		DbConfig()["test"].(map[interface{}]interface{})["encoding"].(string),
	)
}
