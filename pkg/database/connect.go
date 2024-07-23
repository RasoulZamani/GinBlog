package database

import (
	"fmt"
	"log"

	"github.com/RasoulZamani/hiGin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	myEnv := config.ReadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", //sslmode=disable
		myEnv["MYSQL_USER"],
		myEnv["MYSQL_PASSWORD"],
		myEnv["MYSQL_HOST"],
		myEnv["MYSQL_PORT"],
		myEnv["MYSQL_DB_NAME"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(dsn)
		log.Fatal("Could not connect to mysql database!")

		return
	}
	log.Println("Connected to mysql database successfully")
	DB = db
}
