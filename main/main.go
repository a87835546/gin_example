package main

import (
	"gin_example/logic"
	"gin_example/routers"
)

func main() {
	app := routers.InitRouter()
	logic.InitDb()
	logic.InitRedis()
	app.Run(":8080") // listen and serve on 0.0.0.0:8080
}
