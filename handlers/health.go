package handlers

import (
	"net/http"

	"bookstore/log"

	"github.com/gorilla/mux"
)

type HealthHandler struct {
	log *log.Factory
}

func NewHealthHandler(log *log.Factory) *LoginHandler {
	return &LoginHandler{
		log: log,
	}
}

func (h *HealthHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/heartbeat", h.heartbeat).Methods(http.MethodGet)
	router.HandleFunc("/ping", h.ping).Methods(http.MethodGet)
}

func (h *HealthHandler) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong}"))
	if err != nil {
		h.log.Error("heartbeat error", err)
	}
}

func (h *HealthHandler) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{}"))
	if err != nil {
		h.log.Error("heartbeat error", err)
	}
}
