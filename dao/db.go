package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

var connection = "root:longerwu1995@/godemo?charset=utf8&parseTime=True&loc=Local"

func NewDB() (db *gorm.DB) {
	if DB == nil {
		if db, err := gorm.Open("mysql", connection);err != nil {
			panic(err)
		} else{
			db.LogMode(true)
			DB = db
		}
	}
	return DB
}