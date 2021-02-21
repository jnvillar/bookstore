package books

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"
)

type Backend interface {
	Create(book *Book) (*Book, error)
	List(page int) ([]*Book, error)
}

type Factory struct {
	backend Backend
	log     *log.Factory
	time    *time.Factory
}

var factory = map[config.BooksBackend]func(backend *config.BooksConfig) Backend{
	config.BooksMemoryBackend: newMemoryBackend,
}

func NewBookFactory(conf *config.BooksConfig, log *log.Factory, time *time.Factory) *Factory {
	return &Factory{
		backend: factory[conf.Backend](conf),
		log:     log,
		time:    time,
	}
}

func (f *Factory) Create(book *Book) (*Book, error) {
	return f.backend.Create(book)
}

func (f *Factory) List(page int) ([]*Book, error) {
	return f.backend.List(page)
}
