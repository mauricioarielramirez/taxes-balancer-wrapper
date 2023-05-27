package server

import (
	"net/http"
	"taxes-balancer-wrapper/pkg/infraestructure/delivery/routes"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}
	rut := routes.NewRouter()
	a.router = rut.Router()
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
