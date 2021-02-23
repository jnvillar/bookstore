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

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	api := router.Group("/api")

	appHandlers := []handlers.Handler{
		handlers.NewLoginHandler(a.log, a.users, a.validator, a.auth),
		handlers.NewHealthHandler(a.log),
		handlers.NewBooksHandler(a.log, a.validator, a.auth, a.books),
	}

	for _, handler := range appHandlers {
		handler.RegisterRoutes(api)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.AppConfig.Port), router); err != nil {
		panic(err)
	}

	router.Use()
}
