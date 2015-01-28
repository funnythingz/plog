package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbmap gorm.DB

func DbConnect() {
	dbmap, _ = gorm.Open("mysql", "hoge:hoge@/gogogo_development?charset=utf8&parseTime=True")

	dbmap.DB()

	dbmap.DB().Ping()
	dbmap.DB().SetMaxIdleConns(10)
	dbmap.DB().SetMaxOpenConns(100)

	dbmap.SingularTable(true)
}
