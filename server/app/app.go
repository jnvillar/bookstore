package app

import (
	"bookstore/analytics"
	"bookstore/auth"
	"bookstore/books"
	"bookstore/config"
	"bookstore/handlers"
	"bookstore/log"
	"bookstore/middlewares"
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
	analytics *analytics.Factory
}

func NewApp(config *config.Config) *App {
	time := time.NewTimeFactory()
	log := log.NewLoggerFactory(config.LogConfig, time)
	books := books.NewBookFactory(config.BooksConfig, log, time)
	validator := validator.NewMessageValidator()
	users := users.NewUserFactory(config.UsersConfig, log, time)
	auth := auth.NewSessionsFactory(config.AuthConfig, log, time)
	analytics := analytics.NewAnalyticsFactory(config.AnalyticsConfig, log, time)

	app := &App{
		log:       log,
		books:     books,
		time:      time,
		config:    config,
		validator: validator,
		users:     users,
		auth:      auth,
		analytics: analytics,
	}

	return app
}

func (a *App) Init() {
	router := gin.Default()

	a.registerMiddleware(router)
	a.registerRoutes(router)
}

func (a *App) registerMiddleware(router *gin.Engine) {

	appMiddleware := []middlewares.Middleware{
		middlewares.NewAnalyticsMiddleware(a.analytics, a.log),
	}

	for _, middleware := range appMiddleware {
		middleware.RegisterMiddleware(router)
	}
}

func (a *App) registerRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowedMethods:  []string{"*"},
		AllowedHeaders:  []string{"*"},
	}))

	api := router.Group("/api")

	appHandlers := []handlers.Handler{
		handlers.NewLoginHandler(a.log, a.users, a.validator, a.auth),
		handlers.NewHealthHandler(a.log),
		handlers.NewBooksHandler(a.log, a.validator, a.auth, a.books),
		handlers.NewAnalyticsHandler(a.log, a.analytics),
	}

	for _, handler := range appHandlers {
		handler.RegisterRoutes(api)
	}

	if err := router.Run(); err != nil {
		panic(err)
	}
}
