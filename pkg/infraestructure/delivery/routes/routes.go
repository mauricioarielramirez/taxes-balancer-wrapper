package routes

import (
	"net/http"
	"taxes-balancer-wrapper/pkg/infraestructure/handlers"

	"github.com/gorilla/mux"
)

type routes struct {
	PingHandler handlers.Handle
}

type Routes interface {
	Router() *mux.Router
}

func NewRouter() Routes {
	return &routes{}
}

func (r *routes) Router() *mux.Router {
	r.loadHandlers()
	return r.setPaths()
}

func (r *routes) setPaths() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/ping", r.PingHandler.Handle).Methods(http.MethodGet)

	return muxRouter
}

func (r *routes) loadHandlers() {
	r.PingHandler = handlers.NewPingHandler()
}
