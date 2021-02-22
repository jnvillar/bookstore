package books

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"bookstore/config"
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
	jsonFile, err := os.Open("./assets/books.json")
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var books []*Book
	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		panic(err)
	}
	return &memoryBackend{
		config: config,
		books:  books,
	}
}

func (m *memoryBackend) List(page int) ([]*Book, error) {
	return m.books, nil
}
