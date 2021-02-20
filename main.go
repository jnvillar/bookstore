package main

import (
	"bookstore/books"
	"bookstore/config"
	"bookstore/log"
)

func main() {
	app := newApp(config.DevConfig())
}

type App struct {
	log   *log.Factory
	books *books.Factory
}

func newApp(config *config.Config) *App {
	log := log.NewLoggerFactory(config.LogConfig)
	books := books.NewBookFactory(config.BooksConfig, log)

	return &App{
		log:   log,
		books: books}
}
