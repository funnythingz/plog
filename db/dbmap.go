package dbmap

import (
	"../helper"
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
	adapter := DbConfig()["development"].(map[interface{}]interface{})["adapter"].(string)
	encoding := DbConfig()["development"].(map[interface{}]interface{})["encoding"].(string)
	username := DbConfig()["development"].(map[interface{}]interface{})["username"].(string)
	password := DbConfig()["development"].(map[interface{}]interface{})["password"].(string)
	database := DbConfig()["development"].(map[interface{}]interface{})["database"].(string)

	DbOpen(adapter, username, password, database, encoding)
}

func DbTestConnect() {
	adapter := DbConfig()["test"].(map[interface{}]interface{})["adapter"].(string)
	encoding := DbConfig()["test"].(map[interface{}]interface{})["encoding"].(string)
	username := DbConfig()["test"].(map[interface{}]interface{})["username"].(string)
	password := DbConfig()["test"].(map[interface{}]interface{})["password"].(string)
	database := DbConfig()["test"].(map[interface{}]interface{})["database"].(string)

	DbOpen(adapter, username, password, database, encoding)
}
