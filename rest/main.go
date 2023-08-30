package main

import (
	"rest/server"
)

const addressServer string = ":8080"

func main() {

	// crear servidor
	srv := server.New(addressServer)
	// funcion bloqueante
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
