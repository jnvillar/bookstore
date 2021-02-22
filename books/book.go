package books

import (
	"time"

	"bookstore/item"
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
			ID: "c9db3355-a317-4bae-bc1e-7bc912a98463",
			CreatedAt: time.Now(),
			Name: "testName",
		},
		Author: author,
	}
}
