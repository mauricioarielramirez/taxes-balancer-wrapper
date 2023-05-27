package handlers

import "net/http"

type Handle interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
