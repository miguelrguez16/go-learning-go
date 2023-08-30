package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w, r)
		case http.MethodPost:
			addCountries(w, r)
		default:
			sendError(w)
		}
	})
}

func sendError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := fmt.Fprintf(w, "\nMethod not allowed\n")
	if err != nil {
		panic(err)
	}
}
