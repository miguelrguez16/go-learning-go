package server

import (
	"net/http"
)

type Country struct {
	name     string
	language string
}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
