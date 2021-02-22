package app

import (
	"fmt"
	"net/http"

	"bookstore/auth"
	"bookstore/books"
	"bookstore/config"
	"bookstore/handlers"
	"bookstore/log"
	"bookstore/time"
	"bookstore/users"
	"bookstore/validator"

	"github.com/gorilla/mux"
)

type App struct {
	time      *time.Factory
	log       *log.Factory
	books     *books.Factory
	config    *config.Config
	users     *users.Factory
	validator *validator.MessageValidator
	auth      *auth.Factory
}

func NewApp(config *config.Config) *App {
	time := time.NewTimeFactory()
	log := log.NewLoggerFactory(config.LogConfig, time)
	books := books.NewBookFactory(config.BooksConfig, log, time)
	validator := validator.NewMessageValidator()
	users := users.NewUserFactory(config.UsersConfig, log, time)
	auth := auth.NewSessionsFactory(config.AuthConfig, log, time)

	app := &App{
		log:       log,
		books:     books,
		time:      time,
		config:    config,
		validator: validator,
		users:     users,
		auth:      auth,
	}

	return app
}

func (a *App) Init() {
	a.registerRoutes()
}

func (a *App) registerRoutes() {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	handlers := []handlers.Handler{
		handlers.NewLoginHandler(a.log, a.users, a.validator, a.auth),
		handlers.NewHealthHandler(a.log),
		handlers.NewBooksHandler(a.log, a.validator, a.auth, a.books),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes(router)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.AppConfig.Port), router); err != nil {
		panic(err)
	}
}
