package server

import (
	"net/http"
)

type Country struct {
	Name     string
	Language string
}

var countries []*Country

// New create a server on address
func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
