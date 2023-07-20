package logic

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	dbUser     string = "root"
	dbPassword string = "123456"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "demo"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

var Db *gorm.DB
var Client *redis.Client

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}
func InitRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	if err := Client.Ping().Err(); err != nil {
		log.Println("err--->>>", err.Error())
	} else {
		//Client.Set("test", "a", -1)
		log.Println("redis connect success")
	}
	return nil
}
func connectDB() *gorm.DB {
	var err error
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	} else {
		log.Println("db connect success")
	}

	return db
}
