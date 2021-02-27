package books

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
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

func loadContent(fileName string) []*Book {
	// heroku
	content, err := ioutil.ReadFile(fmt.Sprintf("./%s.json", fileName))
	if err != nil {
		// local
		content, err = ioutil.ReadFile(fmt.Sprintf("./server/scrapper/%s.json", fileName))
		if err != nil {
			panic(err)
		}
	}

	var books []*Book
	err = json.Unmarshal(content, &books)

	var res []*Book
	for _, book := range books {
		if book.PictureURL != "" {
			res = append(res, book)
		}
	}

	return res
}

func newMemoryBackend(config *config.BooksConfig) Backend {

	books := loadContent("meli")
	books = append(books, loadContent("distribuidoralabotica")...)

	for _, book := range books {
		book.ID = uuid.New().String()
		book.Price = book.Price * 3
	}
	return &memoryBackend{
		config: config,
		books:  books,
	}
}

func (m *memoryBackend) List(bookSearch *BookSearch) ([]*Book, error) {
	var res []*Book

	if bookSearch.GetName() == "" {
		res = make([]*Book, len(m.books))
		copy(res, m.books)
	} else {
		res = make([]*Book, 0)
		for _, book := range m.books {
			if strings.Contains(strings.ToLower(book.Name), strings.ToLower(bookSearch.Name)) ||
				book.HasAuthor(strings.ToLower(bookSearch.Name)) {
				res = append(res, book)
			}
		}
	}

	if m.config.PageSize < len(res) {
		res = res[0:m.config.PageSize]
	}

	switch bookSearch.PriceOrder {
	case DESC:
		sort.Slice(res[:], func(i, j int) bool {
			return res[i].Price > res[j].Price
		})
	case ASC:
		sort.Slice(res[:], func(i, j int) bool {
			return res[i].Price < res[j].Price
		})

	}

	return res, nil
}
