package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bookstore/books"
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"
	"bookstore/users"

	"github.com/gorilla/mux"
)

type App struct {
	time      *time.Factory
	log       *log.Factory
	books     *books.Factory
	config    *config.Config
	users     *users.Factory
	validator *MessageValidator
}

func NewApp(config *config.Config) *App {
	time := time.NewTimeFactory()
	log := log.NewLoggerFactory(config.LogConfig, time)
	books := books.NewBookFactory(config.BooksConfig, log, time)
	validator := NewMessageValidator()
	users := users.NewUserFactory(config.UsersConfig, log, time)

	app := &App{
		log:       log,
		books:     books,
		time:      time,
		config:    config,
		validator: validator,
		users:     users,
	}

	return app
}

func (a *App) Init() {
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	router := mux.NewRouter()

	// health
	router.HandleFunc("/heartbeat", a.heartbeat).Methods(http.MethodGet)
	router.HandleFunc("/ping", a.ping).Methods(http.MethodGet)

	// auth
	router.HandleFunc("/login", a.login).Methods(http.MethodPost)
	router.HandleFunc("/logout", a.logout).Methods(http.MethodGet)

	//books
	router.HandleFunc("/books", a.listBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", a.createBook).Methods(http.MethodPost)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.AppConfig.Port), router); err != nil {
		panic(err)
	}
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func (a *App) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func (a *App) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{}"))
	if err != nil {
		a.log.Error("ping error", err)
	}
}

func (a *App) listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := a.validator.ValidateListBooks(r); err != nil {
		a.log.Error("error validating list books request", err)
		writeError(http.StatusBadRequest, w, err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		a.log.Error("error parsing page number", err)
		page = 0
	}

	books, err := a.books.List(page)
	if err != nil {
		a.log.Error("error listing books", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		a.log.Error("error marshalling list books", err)
		writeError(http.StatusInternalServerError, w, err)
		return
	}
}

func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	book := &books.Book{}
	err := decoder.Decode(book)
	if err != nil {
		writeError(http.StatusBadRequest, w, err)
		return
	}

	if err := a.validator.ValidateCreateBook(book); err != nil {
		a.log.Error("error validating create book request", err)
		writeError(http.StatusBadRequest, w, err)
		return
	}

	book, err = a.books.Create(book)
	if err != nil {
		a.log.Error("error creating book", err)
		writeError(http.StatusInternalServerError, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		a.log.Error("error marshalling book", err)
		writeError(http.StatusInternalServerError, w, err)
		return
	}
}
