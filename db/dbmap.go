package db

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/k0kubun/pp"
	"log"
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

type env struct {
	Database string
}

type connection struct {
	Adapter  string
	Encoding string
	Username string
	Password string
}

type Config struct {
	Connection connection
	Databases  map[string]env
}

func DbConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("database.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}

func DbConnect(env string) {
	config := DbConfig()

	DbOpen(
		config.Connection.Adapter,
		config.Connection.Username,
		config.Connection.Password,
		config.Databases[env].Database,
		config.Connection.Encoding,
	)
}
