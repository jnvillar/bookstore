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

	"github.com/gorilla/mux"
)

type App struct {
	time   *time.Factory
	log    *log.Factory
	books  *books.Factory
	config *config.Config
}

func NewApp(config *config.Config) *App {
	time := time.NewTimeFactory()
	log := log.NewLoggerFactory(config.LogConfig, time)
	books := books.NewBookFactory(config.BooksConfig, log, time)

	app := &App{
		log:    log,
		books:  books,
		time:   time,
		config: config,
	}

	return app
}

func (a *App) Init() {
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/heartbeat", a.heartbeat).Methods(http.MethodGet)
	router.HandleFunc("/ping", a.ping).Methods(http.MethodGet)

	router.HandleFunc("/books", a.listBooks).Methods(http.MethodGet)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.AppConfig.Port), router); err != nil {
		panic(err)
	}
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
