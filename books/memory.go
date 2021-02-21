package books

import (
	"bookstore/config"
)

type memoryBackend struct {
	books  []*Book
	config *config.BooksConfig
}

func (m *memoryBackend) Create(book *Book) (*Book, error) {
	m.books = append(m.books, book)
	return book, nil
}

func newMemoryBackend(config *config.BooksConfig) Backend {
	return &memoryBackend{
		config: config,
		books: []*Book{
			newBook("juanito"),
		},
	}
}

func (m *memoryBackend) List(page int) ([]*Book, error) {
	return m.books, nil
}
