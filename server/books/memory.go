package books

import (
	"encoding/json"
	"errors"
	"strings"

	"bookstore/assets"
	"bookstore/config"

	"github.com/google/uuid"
)

type memoryBackend struct {
	books  []*Book
	config *config.BooksConfig
}

func (m *memoryBackend) Get(bookID string) (*Book, error) {
	for _, book := range m.books {
		if book.ID == bookID {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (m *memoryBackend) Update(book *Book) (*Book, error) {
	oldBook, err := m.Get(book.ID)
	if err != nil {
		return nil, err
	}
	oldBook.Name = book.Name
	oldBook.Author = book.Author
	oldBook.Stock = book.Stock
	oldBook.PictureURL = book.PictureURL
	oldBook.Price = book.Price
	return oldBook, err
}

func (m *memoryBackend) Create(book *Book) (*Book, error) {
	m.books = append(m.books, book)
	return book, nil
}

func newMemoryBackend(config *config.BooksConfig) Backend {
	var books []*Book
	err := json.Unmarshal(assets.Books, &books)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		book.Name = strings.Title(strings.ToLower(book.Name))
		book.ID = uuid.New().String()
		book.Price = book.Price * 3
	}
	return &memoryBackend{
		config: config,
		books:  books,
	}
}

func (m *memoryBackend) List(page int) ([]*Book, error) {
	return m.books, nil
}
