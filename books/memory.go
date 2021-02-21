package books

import (
	"bookstore/config"
	"bookstore/item"
)

type memoryBackend struct {
	books  []*Book
	config *config.BooksConfig
}

func newMemoryBackend(config *config.BooksConfig) Backend {
	return &memoryBackend{
		config: config,
		books: []*Book{
			{
				Item: &item.Item{
					ID:         "0",
					Price:      0,
					Name:       "",
					PictureURL: "",
					Stock:      0,
				},
			},
		},
	}
}

func (m *memoryBackend) List(page int) ([]*Book, error) {
	return m.books, nil
}
