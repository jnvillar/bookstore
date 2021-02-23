package handlers

import "github.com/gorilla/mux"

type Handler interface {
	RegisterRoutes(router *mux.Router)
}
