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

	"github.com/gorilla/mux"
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

func (b *BooksHandler) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/books").Subrouter()
	subRouter.HandleFunc("", b.listBooks).Methods(http.MethodGet)
	subRouter.HandleFunc("", b.createBook).Methods(http.MethodPost)
	subRouter.HandleFunc("", b.updateBook).Methods(http.MethodPut)
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

	func (b *BooksHandler) listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		b.log.Error("error parsing page number", err)
		page = 0
	}

	books, err := b.books.List(page)
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
