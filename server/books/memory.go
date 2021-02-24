package books

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

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
	content, err := ioutil.ReadFile("./server/scrapper/output.json")
	if err != nil {
		panic(err)
	}
	var books []*Book
	err = json.Unmarshal(content, &books)
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

func (m *memoryBackend) List(bookSearch *BookSearch, page int) ([]*Book, error) {
	res := make([]*Book, 0)
	if bookSearch.GetName() == "" {
		if m.config.PageSize < len(m.books) {
			return m.books[0:m.config.PageSize], nil
		}
		return m.books, nil
	}
	for _, book := range m.books {
		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(bookSearch.Name)) {
			res = append(res, book)
		}
	}
	return res, nil
}
