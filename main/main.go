package main

import (
	"fmt"
	"gin_example/logic"
	"gin_example/routers"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	app := routers.InitRouter()
	logic.InitDb()
	err := logic.InitRedis()
	if err != nil {
		return
	}
	err = app.Run(":8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080

}
