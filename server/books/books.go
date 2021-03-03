package books

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type Backend interface {
	Get(id string) (*Book, error)
	Update(book *Book) (*Book, error)
	Create(book *Book) (*Book, error)
	List(search *BookSearch) ([]*Book, error)
	Visit(bookID string) error
	GetCategories() ([]string, error)
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
	book.ID = uuid.New().String()
	book.CreatedAt = f.time.Now()
	return f.backend.Create(book)
}

func (f *Factory) List(search *BookSearch) ([]*Book, error) {
	return f.backend.List(search)
}

func (f *Factory) GetCategories() ([]string, error) {
	return f.backend.GetCategories()
}

func (f *Factory) Get(bookID string) (*Book, error) {
	return f.backend.Get(bookID)
}

func (f *Factory) Visit(bookID string) error {
	return f.backend.Visit(bookID)
}

func (f *Factory) Update(newBook map[string]interface{}) (*Book, error) {
	book, err := f.Get(newBook["id"].(string))
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(newBook, book)
	if err != nil {
		return nil, err
	}

	return f.backend.Update(book)
}
