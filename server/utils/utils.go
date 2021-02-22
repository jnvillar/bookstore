package utils

import (
	"net/http"
)

func WriteError(status int, w http.ResponseWriter, err error) {
	w.WriteHeader(status)
	// nolint:errcheck
	w.Write([]byte(err.Error()))
}

func WriteOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	// nolint:errcheck
	w.Write([]byte("Ok"))
}
