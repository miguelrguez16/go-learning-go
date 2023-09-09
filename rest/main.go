package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rest/server"
	"syscall"
)

const addressServer string = ":8080"

func main() {

	ctx := context.Background()

	serverDoneChannel := make(chan os.Signal, 1)
	signal.Notify(serverDoneChannel, os.Interrupt, syscall.SIGTERM)

	srv := server.New(addressServer)

	// start server
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Panicln(err)
		}
	}()
	log.Println("Server started")

	<-serverDoneChannel

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Server end")

}
