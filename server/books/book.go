package books

import (
	"strings"
	"time"

	"bookstore/item"
)

type Book struct {
	*item.Item `mapstructure:",squash"`
	Author     []string `json:"author"`
	Category   []string `json:"category"`
	Publisher  []string `json:"publisher"`
}

func (b *Book) GetAuthor() []string {
	if b != nil {
		return b.Author
	}
	return []string{}
}

func (b *Book) HasCategory(search string) bool {
	if b == nil {
		return false
	}
	for _, cat := range b.Category {
		return strings.Contains(strings.ToLower(cat), strings.ToLower(search))
	}
	return false
}

func (b *Book) HasAuthor(search string) bool {
	if b == nil {
		return false
	}
	for _, author := range b.Author {
		return strings.Contains(strings.ToLower(author), strings.ToLower(search))
	}
	return false
}

func newBook(author string) *Book {
	return &Book{
		Item: &item.Item{
			ID:        "c9db3355-a317-4bae-bc1e-7bc912a98463",
			CreatedAt: time.Now(),
			Name:      "testName",
		},
		Author: []string{author},
	}
}
