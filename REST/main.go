package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	/**
	http.ResponseWriter -> enviar repuesta
	*http.Request -> donde va la peticion
	*/
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprintf(writer, "\nMethod not allowed\n")
			if err != nil {
				panic(err)
			}
		} else {
			writer.WriteHeader(http.StatusOK)
			_, err := fmt.Fprintf(writer, "\nHello there %v\n", time.Now().Nanosecond())
			if err != nil {
				panic(err)
			}
		}
	})

	// crear servidor
	srv := http.Server{
		Addr: ":8080",
	}
	// funcion bloqueante
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
