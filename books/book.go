package books

import (
	"bookstore/item"

	"github.com/google/uuid"
)

type Book struct {
	*item.Item
	Author string `json:"author"`
}

func (b *Book) GetAuthor() string {
	if b != nil {
		return b.Author
	}
	return ""
}

func newBook(author string) *Book {
	return &Book{
		Item: &item.Item{
			ID: uuid.New().String(),
		},
		Author: author,
	}
}
