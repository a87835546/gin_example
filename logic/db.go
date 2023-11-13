package logic

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/go-redis/redis"
	lua "github.com/yuin/gopher-lua"
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
	ruleBbName string = "rule_db"
)

var (
	dsn     = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	RuleDSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, ruleBbName)
	Db      *gorm.DB
	RuleDb  *gorm.DB
	Client  *redis.Client
	E       *casbin.Enforcer
)

func InitDb() *gorm.DB {
	Db = connectDB()
	RuleDb = connectRuleDB()
	return Db
}

var script = `
		local value = redis.call("Get", KEYS[1])
		  local result  = ''
			local arg={...}
			for i,v in ipairs(arg) do
				result = result .. v
			end
		return value
`
var script1 = `
	local key = KEYS[1]
	local sum = redis.call("get",key)
	if not sum then
	  sum = 0
	end
	local num_arg = #ARGV
	for i =1,num_arg do
		sum = sum + ARGV[i]
	end
	redis.call("set",key,sum)
	return sum
`

func useLua() {
	//编写脚本 - 检查数值，是否够用，够用再减，否则返回减掉后的结果
	var luaScript = redis.NewScript(`
		local value = redis.call("Get", KEYS[1])
		return value
	`)
	//执行脚本
	n, err := luaScript.Run(Client, []string{"test", "6"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
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
		log.Println("redis connect success")

		useLua()
		l := redis.NewScript(script1)
		if n, err := l.Run(Client, []string{"test1"}, 1, 2, 3).Result(); err != nil {
			log.Println("get redis value err--->>>", err.Error())
		} else {
			log.Printf("redis get value --->>> %s\n", n)
		}
		InitLua()
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
func connectRuleDB() *gorm.DB {
	var err error
	fmt.Println("dsn : ", RuleDSN)
	db, err := gorm.Open(mysql.Open(RuleDSN), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

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
func InitLua() *lua.LState {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile("controllers/redis.lua"); err != nil {
		panic(err)
	}
	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("getScriptValue"),
		NRet:    1,
		Protect: true,
	}, lua.LString("test"), lua.LString("test1"), lua.LString("test2"))
	if err != nil {
		panic(err)
	}
	ret := l.Get(-1)
	l.Pop(1)
	res, ok := ret.(lua.LString)
	if ok {
		log.Printf("res-->>%s", res)
	} else {
		log.Printf("err")
	}
	return l
}
