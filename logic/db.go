package logic

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	dbUser string = "root"
	//dbPassword string = "insert_password"
	dbPassword string = "12345678"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "dev_db"
)

var (
	dsn    = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	Db     *gorm.DB
	Client *redis.Client
	E      *casbin.Enforcer
)

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
func InitCasbin() {
	a, err := gormadapter.NewAdapter("mysql", "root:12345678@tcp(127.0.0.1:3306)/dev_db", true)
	if err != nil {
		log.Fatalf("mysql failed:%v\n", err)
	}
	e, err := casbin.NewEnforcer("model/model.conf", a)
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	check(e, "dajun", "data", "read")
	check(e, "alice", "data", "read")
	E = e
}
func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
