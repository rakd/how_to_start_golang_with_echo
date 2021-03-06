package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var isConnected bool

// Setup ..
func Setup() {
	if isConnected == true {
		return
	}
	log.Print("models:init")

	dbHost := getenvWithDefault("DATABASE_HOST", "localhost")
	dbName := getenvWithDefault("DATABASE_NAME", "sample")
	dbUser := getenvWithDefault("DATABASE_USER", "root")
	dbPass := getenvWithDefault("DATABASE_PASS", "")
	dbPort := getenvWithDefault("DATABASE_PORT", "3306")
	if dbPass != "" {
		dbPass = ":" + dbPass
	}

	dsn1 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	dsn2 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		"mysql",
	)

	var dberr error

	loop := 0
	for true {
		db, dberr = gorm.Open("mysql", dsn1)

		log.Printf("dns1:%s", dsn1)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn1)
			isConnected = true
			break
		}
		log.Print("DB Connection Error: dsn1=" + dsn1)
		log.Print(dberr.Error())

		time.Sleep(time.Millisecond * 3000)

		db, dberr = gorm.Open("mysql", dsn2)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn2)
			log.Printf("create database %s", dbName)
			db.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			// might be error...?
			log.Printf("use dbname")
			db.Exec("use `" + dbName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			db.Close()
		} else {
			log.Print("DB Connection Error: dsn2=" + dsn2)
			log.Print(dberr.Error())
			loop++
			if loop > 300 {
				//give up to connect... even though no connect db.
				break
			}
		}

	}

	log.Printf("before auto migrate.")
	MigrateTables()
	log.Printf("after auto migrate..")
}

// getenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

// MigrateTables ...
func MigrateTables() {

	log.Print("User")
	db.AutoMigrate(&User{})

}
