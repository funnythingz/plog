package db

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/k0kubun/pp"
	"log"
)

var (
	Dbmap gorm.DB
)

func DbOpen(
	adapter string,
	host string,
	username string,
	password string,
	database string,
	encoding string,
) {

	dataSourceName := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=" + encoding + "&parseTime=True"
	Dbmap, _ = gorm.Open(adapter, dataSourceName)

	Dbmap.DB()

	Dbmap.DB().Ping()
	Dbmap.DB().SetMaxIdleConns(10)
	Dbmap.DB().SetMaxOpenConns(100)

	Dbmap.SingularTable(true)
}

type env struct {
	Host     string
	Username string
	Password string
	Database string
}

type connection struct {
	Adapter  string
	Encoding string
	Host     string
	Username string
	Password string
	Database string
}

type ConfigData struct {
	Connection connection
	Databases  map[string]env
}

func DbConfig() ConfigData {
	var config ConfigData
	if _, err := toml.DecodeFile("database.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}

var (
	Config = DbConfig()
)

func DbConnect(env string) {
	var (
		dbEnv = Config.Databases[env]

		host = func() string {
			if dbEnv.Host != "" {
				return dbEnv.Host
			}
			return Config.Connection.Host
		}()

		username = func() string {
			if dbEnv.Username != "" {
				return dbEnv.Username
			}
			return Config.Connection.Username
		}()

		password = func() string {
			if dbEnv.Password != "" {
				return dbEnv.Password
			}
			return Config.Connection.Password
		}()

		database = func() string {
			if dbEnv.Database != "" {
				return dbEnv.Database
			}
			return Config.Connection.Database
		}()
	)

	DbOpen(
		Config.Connection.Adapter,
		host,
		username,
		password,
		database,
		Config.Connection.Encoding,
	)
}
