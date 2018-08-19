package db

import (
	"fmt"
	"go-crawler/search-api/api/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "root")
	dbname := getEnv("DB_NAME", "go_crawler")

	dbinfo := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True",
		user,
		password,
		dbname,
	)
	println(dbinfo)
	db, err = gorm.Open("mysql", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.AnimeDocument{}) {
		err := db.CreateTable(&models.AnimeDocument{})
		if err != nil {
			log.Println("Table already exists")
		}
	}
	db.AutoMigrate(&models.AnimeDocument{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
