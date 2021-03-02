package handlers

import (
	"net/http"

	"bookstore/log"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	log *log.Factory
}

func NewHealthHandler(log *log.Factory) *HealthHandler {
	return &HealthHandler{
		log: log,
	}
}

func (h *HealthHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/heartbeat", func(c *gin.Context) { h.heartbeat(c.Writer, c.Request) })
	router.GET("/ping", func(c *gin.Context) { h.ping(c.Writer, c.Request) })

	//router.HandleFunc("/heartbeat", h.heartbeat).Methods(http.MethodGet)
	//router.HandleFunc("/ping", h.ping).Methods(http.MethodGet)
}

func (h *HealthHandler) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
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
