package server

import (
	"fmt"
	"net/http"
)

const indexRoute string = "/"
const countriesRoute string = "/countries"

// initRoutes start all routes
func initRoutes() {
	http.HandleFunc(indexRoute, index)
	http.HandleFunc(countriesRoute, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w)
		case http.MethodPost:
			addCountries(w, r)
		default:
			sendError(w)
		}
	})
}

// sendError return message error
func sendError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := fmt.Fprintf(w, "\nMethod not allowed\n")
	if err != nil {
		panic(err)
	}
}
