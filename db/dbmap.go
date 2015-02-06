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

func DbOpen(adapter string,
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

func DbDevelopmentConnect() {
	m := make(map[interface{}]interface{})
	config, _ := ioutil.ReadFile("database.yml")
	err := yaml.Unmarshal([]byte(config), &m)
	helper.Check(err)

	adapter := m["development"].(map[interface{}]interface{})["adapter"].(string)
	encoding := m["development"].(map[interface{}]interface{})["encoding"].(string)
	username := m["development"].(map[interface{}]interface{})["username"].(string)
	password := m["development"].(map[interface{}]interface{})["password"].(string)
	database := m["development"].(map[interface{}]interface{})["database"].(string)

	DbOpen(adapter, username, password, database, encoding)
}

func DbTestConnect() {
	m := make(map[interface{}]interface{})
	config, _ := ioutil.ReadFile("database.yml")
	err := yaml.Unmarshal([]byte(config), &m)
	helper.Check(err)

	adapter := m["test"].(map[interface{}]interface{})["adapter"].(string)
	encoding := m["test"].(map[interface{}]interface{})["encoding"].(string)
	username := m["test"].(map[interface{}]interface{})["username"].(string)
	password := m["test"].(map[interface{}]interface{})["password"].(string)
	database := m["test"].(map[interface{}]interface{})["database"].(string)

	DbOpen(adapter, username, password, database, encoding)
}
