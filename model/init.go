package model

import (
	// _ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db gorm.DB
)

func InitDb() (err error) {
	// db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err := gorm.Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True")
	db, err = gorm.Open("sqlite3", "./gorm.db")

	db.AutoMigrate(User{})

	user := User{Name: "Jijesh", Email: "jijesh@example.com"}
	db.Create(&user)
	return
}
