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

func (m *memoryBackend) Visit(bookID string) error {
	_, err := m.Get(bookID)
	if err != nil {
		return err
	}
	for _, book := range m.books {
		if book.ID == bookID {
			book.Visits += 1
		}
	}
	return nil
}

func (m *memoryBackend) GetCategories() ([]string, error) {
	categories := map[string]string{}
	for _, book := range m.books {
		for _, c := range book.Category {
			categories[c] = c
		}
	}
	res := make([]string, 0)
	for k, _ := range categories {
		res = append(res, k)
	}
	return res, nil
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
	_, err := m.Get(book.ID)
	if err != nil {
		return nil, err
	}
	res := make([]*Book, 0)
	for _, b := range m.books {
		if b.ID == book.ID {
			res = append(res, book)
		} else {
			res = append(res, b)
		}
	}

	return book, nil
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

	meliBooks := loadContent("meli")

	for _, book := range meliBooks {
		book.ID = uuid.New().String()
		book.Price = int64(float64(book.Price) * 2)
	}

	botanicaBooks := loadContent("distribuidoralabotica")

	for _, book := range botanicaBooks {
		book.ID = uuid.New().String()
		book.Price = int64(float64(book.Price) * 2.5)
	}

	return &memoryBackend{
		config: config,
		books:  append(meliBooks, botanicaBooks...),
	}
}

func (m *memoryBackend) List(bookSearch *BookSearch) ([]*Book, error) {
	booksCopy := make([]*Book, len(m.books))
	copy(booksCopy, m.books)

	// sort by visits by default
	sort.Slice(booksCopy[:], func(i, j int) bool {
		return booksCopy[i].Visits > booksCopy[j].Visits
	})

	filteredByCategory := make([]*Book, 0)
	if bookSearch.Category != "" {
		for _, book := range booksCopy {
			if book.HasCategory(bookSearch.Category) {
				filteredByCategory = append(filteredByCategory, book)
			}
		}
		booksCopy = filteredByCategory
	}

	// sort by price
	switch bookSearch.PriceOrder {
	case DESC:
		sort.Slice(booksCopy[:], func(i, j int) bool {
			return booksCopy[i].Price > booksCopy[j].Price
		})
	case ASC:
		sort.Slice(booksCopy[:], func(i, j int) bool {
			return booksCopy[i].Price < booksCopy[j].Price
		})
	}

	// filter by name
	filteredByName := make([]*Book, 0)
	if bookSearch.GetName() != "" {
		for _, book := range booksCopy {
			if strings.Contains(strings.ToLower(book.Name), strings.ToLower(bookSearch.Name)) ||
				book.HasAuthor(strings.ToLower(bookSearch.Name)) {
				filteredByName = append(filteredByName, book)
			}
		}
		booksCopy = filteredByName
	}

	// apply pagination
	start := bookSearch.Page * m.config.PageSize
	if len(booksCopy) > start {
		booksCopy = booksCopy[start:]
	} else {
		booksCopy = []*Book{}
	}

	// trim by pagesize
	if len(booksCopy) > m.config.PageSize {
		booksCopy = booksCopy[0:m.config.PageSize]
	}

	return booksCopy, nil
}
