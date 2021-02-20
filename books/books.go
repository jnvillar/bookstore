package books

import (
	"bookstore/config"
	"bookstore/log"
)

type Backend interface {
}

type Factory struct {
	log *log.Factory
}

func NewBookFactory(conf *config.BooksConfig, log *log.Factory) *Factory {
	return &Factory{log: log}
}
