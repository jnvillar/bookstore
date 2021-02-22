package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/auth"
	"bookstore/log"
	"bookstore/login"
	"bookstore/users"
	"bookstore/utils"
	validator2 "bookstore/validator"

	"github.com/gorilla/mux"
)

type LoginHandler struct {
	log       *log.Factory
	validator *validator2.MessageValidator
	users     *users.Factory
	auth      *auth.Factory
}

func NewLoginHandler(log *log.Factory, users *users.Factory, validator *validator2.MessageValidator, auth *auth.Factory) *LoginHandler {
	return &LoginHandler{
		log:       log,
		users:     users,
		validator: validator,
		auth:      auth,
	}
}

func (l *LoginHandler) RegisterRoutes(router *mux.Router) {
	// auth
	router.HandleFunc("/login", l.login).Methods(http.MethodPost)
	router.HandleFunc("/logout", l.logout).Methods(http.MethodGet)
}

func (l *LoginHandler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	loginRequest := &login.Request{}
	err := decoder.Decode(loginRequest)
	if err != nil {
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	if err := l.validator.ValidateLogin(loginRequest); err != nil {
		l.log.Error("error validating create book request", err)
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	user, err := l.users.Get(loginRequest.Username, loginRequest.Password)
	if err != nil {
		l.log.Error("error fetching user", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	session, err := l.auth.Create(user)
	if err != nil {
		l.log.Error("error creating session for user", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(session)
	if err != nil {
		l.log.Error("error marshalling session", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}

func (l *LoginHandler) logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	token := r.Header.Get(auth.AuthHeader)
	if err := l.auth.Delete(token); err != nil {
		l.log.Error("error loging out", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
	}
	utils.WriteOk(w)
}
