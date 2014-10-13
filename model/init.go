package model

import (
	// _ "github.com/go-sql-driver/mysql"

	"log"
	"sync"

	"github.com/jinzhu/gorm"
	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db   gorm.DB
	once sync.Once
)

func InitDb() (err error) {
	// db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err := gorm.Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True")
	once.Do(func() {
		log.Println("Initialize database")
		db, err = gorm.Open("sqlite3", "./gorm.db")
		db.AutoMigrate(User{})
	})
	return
}
