package app

import (
	"net/http"
)

func writeError(status int ,w http.ResponseWriter, err error){
	w.WriteHeader(status)
	// nolint:errcheck
	w.Write([]byte(err.Error()))
}