package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "log"
	_ "reflect"
)

var dbmap gorm.DB

func DbConnect() {
	m := make(map[interface{}]interface{})
	config, _ := ioutil.ReadFile("database.yml")
	err := yaml.Unmarshal([]byte(config), &m)
	check(err)

	adapter := m["development"].(map[interface{}]interface{})["adapter"].(string)
	encoding := m["development"].(map[interface{}]interface{})["encoding"].(string)
	username := m["development"].(map[interface{}]interface{})["username"].(string)
	password := m["development"].(map[interface{}]interface{})["password"].(string)
	database := m["development"].(map[interface{}]interface{})["database"].(string)

	dbmap, _ = gorm.Open(adapter, username+":"+password+"@/"+database+"?charset="+encoding+"&parseTime=True")

	dbmap.DB()

	dbmap.DB().Ping()
	dbmap.DB().SetMaxIdleConns(10)
	dbmap.DB().SetMaxOpenConns(100)

	dbmap.SingularTable(true)
}
