package main

import (
	"context"
	"errors"
	"fmt"
	"gin_example/logic"
	"gin_example/routers"
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
