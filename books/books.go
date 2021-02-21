package books

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"
)

type Backend interface {
	List(page int) ([]*Book, error)
}

type Factory struct {
	Backend
	log  *log.Factory
	time *time.Factory
}

var factory = map[config.BooksBackend]func(backend *config.BooksConfig) Backend{
	config.BooksMemoryBackend: newMemoryBackend,
}

func NewBookFactory(conf *config.BooksConfig, log *log.Factory, time *time.Factory) *Factory {
	return &Factory{
		Backend: factory[conf.Backend](conf),
		log:     log,
		time:    time,
	}
}
