package main

import (
	"context"
	"errors"
	"fmt"
	"gin_example/logic"
	"gin_example/routers"
	lua "github.com/yuin/gopher-lua"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	loadValue()
	logic.InitDb()
	err := logic.InitRedis()
	logic.InitCasbin()
	if err != nil {
		return
	}
	app := routers.InitRouter()
	err = app.Run(":8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
}

func loadValue() {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile("controllers/redis.lua"); err != nil {
		panic(err)
	}
	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("test1"),
		NRet:    1,
		Protect: true,
	}, lua.LString("123"))
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
}
