package db

import (
	"../helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
	"github.com/jinzhu/gorm"
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
	config := DbConfig()["development"].(map[interface{}]interface{})
	DbOpen(
		config["adapter"].(string),
		config["username"].(string),
		config["password"].(string),
		config["database"].(string),
		config["encoding"].(string),
	)
}

func DbTestConnect() {
	config := DbConfig()["test"].(map[interface{}]interface{})
	DbOpen(
		config["adapter"].(string),
		config["username"].(string),
		config["password"].(string),
		config["database"].(string),
		config["encoding"].(string),
	)
}
