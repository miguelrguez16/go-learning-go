package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"rest/configuration"
	"rest/server"
	"syscall"
)

const addressServer string = ":8080"

//const envPort string = "API_PORT"

func main() {

	ctx := context.Background()
	// db url complete
	url, port, name := getDB()
	fmt.Printf("URL [%v] PORT [%v] NAME [%v]\n", url, port, name)

	// load data from yaml
	load, err := configuration.Load("config.yml")
	if err != nil {
		return
	}
	configuration.PrintConfiguration(load)
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

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Server end")

}

func getDB() (string, string, string) {
	// load file env
	err := godotenv.Load(".env")
	if err != nil {
		return "", "", ""
	}
	url, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Println("DB_URL not found")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		log.Println("DB_PORT not found")
	}
	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Println("DB_NAME not found")
	}

	return url, port, name
}

// envVarsExtract()
/* func envVarsExtract() {
	port, ok := os.LookupEnv(envPort)
	if !ok {
		log.Panicln(ok)
	}
	fmt.Println(port)
}*/
