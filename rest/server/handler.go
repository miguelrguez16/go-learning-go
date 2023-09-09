package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, "\nMethod not allowed\n")
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "\nHello there %v\n", time.Now().Nanosecond())
		if err != nil {
			panic(err)
		}
	}
}

func getCountries(w http.ResponseWriter) {
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(countries)
	if err != nil {
		panic(err)
	}
}

func addCountries(w http.ResponseWriter, r *http.Request) {
	country := &Country{} // tiene q
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "%v", err)
		if err != nil {
			panic(err)
		}
	}
	countries = append(countries, country)
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Added country %s %s", country.Name, country.Language)
	if err != nil {
		panic(err)
	}
}
