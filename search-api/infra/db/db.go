package db

import (
	"fmt"
	repository "go-crawler/search-api/api/repositorys/repository-gorm"
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

	//dbinfo := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True",
	dbinfo := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True",
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

	if !db.HasTable(&repository.AnimeDocument{}) {
		err := db.CreateTable(&repository.AnimeDocument{})
		if err != nil {
			log.Println("Table already exists")
			//db.DropTable(&repository.AnimeDocument{})
		}
	}
	db.AutoMigrate(&repository.AnimeDocument{})
	db.DB().SetMaxIdleConns(22)
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
