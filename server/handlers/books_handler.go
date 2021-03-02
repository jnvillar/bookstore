package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/auth"
	"bookstore/books"
	"bookstore/log"
	"bookstore/permissions"
	"bookstore/utils"
	validator2 "bookstore/validator"

	"github.com/gin-gonic/gin"
)

type BooksHandler struct {
	log       *log.Factory
	validator *validator2.MessageValidator
	auth      *auth.Factory
	books     *books.Factory
}

func NewBooksHandler(log *log.Factory, validator *validator2.MessageValidator, auth *auth.Factory, books *books.Factory) *BooksHandler {
	return &BooksHandler{
		log:       log,
		validator: validator,
		auth:      auth,
		books:     books,
	}
}

func (b *BooksHandler) RegisterRoutes(router *gin.RouterGroup) {
	booksApi := router.Group("/books")
	booksApi.GET("/categories", func(c *gin.Context) { b.listCategories(c.Writer, c.Request) })
	booksApi.GET("", func(c *gin.Context) { b.listBooks(c.Writer, c.Request) })
	booksApi.POST("", func(c *gin.Context) { b.createBook(c.Writer, c.Request) })
	booksApi.PUT("", func(c *gin.Context) { b.updateBook(c.Writer, c.Request) })
}

func (b *BooksHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var book map[string]interface{}
	err := decoder.Decode(&book)
	if err != nil {
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	if err := b.validator.ValidateUpdateBook(book); err != nil {
		b.log.Error("error validating list books request", err)
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	updatedBook, err := b.books.Update(book)
	if err != nil {
		b.log.Error("error updating book", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		b.log.Error("error marshalling updated book", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}

func (b *BooksHandler) listCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := b.books.GetCategories()
	if err != nil {
		b.log.Error("error fetching books categories", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		b.log.Error("error marshalling categories", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}

func (b *BooksHandler) listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		b.log.Error("error parsing page number", err)
		page = 0
	}

	cat := r.URL.Query().Get("cat")
	search := r.URL.Query().Get("search")
	priceOrder := r.URL.Query().Get("price")

	bookSearch := &books.BookSearch{
		Name:       search,
		Page:       page,
		Category:   cat,
		PriceOrder: books.StringToPriceOrder(priceOrder),
	}

	books, err := b.books.List(bookSearch)
	if err != nil {
		b.log.Error("error listing books", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		b.log.Error("error marshalling list books", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}

func (b *BooksHandler) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := b.auth.Validate(r.Header.Get(auth.AuthHeader), permissions.WriteBook)
	if err != nil {
		utils.WriteError(http.StatusUnauthorized, w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	book := &books.Book{}
	err = decoder.Decode(book)
	if err != nil {
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	if err := b.validator.ValidateCreateBook(book); err != nil {
		b.log.Error("error validating create book request", err)
		utils.WriteError(http.StatusBadRequest, w, err)
		return
	}

	book, err = b.books.Create(book)
	if err != nil {
		b.log.Error("error creating book", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		b.log.Error("error marshalling book", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}
