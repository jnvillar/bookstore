package books

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"

	"github.com/google/uuid"
)

type Backend interface {
	Get(id string) (*Book, error)
	Update(book *Book) (*Book, error)
	Create(book *Book) (*Book, error)
	List(search *BookSearch) ([]*Book, error)
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

func (f *Factory) Get(bookID string) (*Book, error) {
	return f.backend.Get(bookID)
}

func (f *Factory) Update(params map[string]interface{}) (*Book, error) {
	id := params["id"].(string)
	book, err := f.Get(id)
	if err != nil{
		return nil, err
	}

	name, ok := params["name"]
	if ok {
		book.Name = name.(string)
	}

	//author, ok := params["author"]
	//if ok {
	//	book.Author = author.(string)
	//}

	price, ok := params["price"]
	if ok {
		book.Price = int64(price.(float64))
	}

	stock, ok := params["stock"]
	if ok {
		book.Stock = int64(stock.(float64))
	}

	pictureUrl, ok := params["pictureUrl"]
	if ok {
		book.PictureURL = pictureUrl.(string)
	}


	now := f.time.Now()
	book.UpdatedAt = &now

	return f.backend.Update(book)
}
