package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect function helps to open a database connection
func Connect() {
	d, err := gorm.Open("mysql", "root:password@123@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

// This GetDB function will used in other files
// It return db
func GetDB() *gorm.DB {
	return db
}
