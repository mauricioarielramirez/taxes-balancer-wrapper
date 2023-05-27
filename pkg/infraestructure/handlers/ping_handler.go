package handlers

import (
	"encoding/json"
	"net/http"
)

type PingHandler struct{}

func NewPingHandler() Handle {
	return &PingHandler{}
}

func (ph *PingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}
