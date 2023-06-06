package logic

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	dbUser     string = "root"
	dbPassword string = "123456abc"
	dbHost     string = "13.215.251.145"
	dbPort     int    = 3306
	dbName     string = "demo"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	} else {
		log.Println("db connect success")
	}

	return db
}
