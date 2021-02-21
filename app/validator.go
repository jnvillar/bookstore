package app

import (
	"net/http"

	"bookstore/books"
	"bookstore/validator"
)

// MessageValidator implements the Validator interface
type MessageValidator struct {
}

func NewMessageValidator() *MessageValidator {
	return &MessageValidator{}
}

func (m *MessageValidator) ValidateListBooks(r *http.Request) error {
	return validator.FirstNonValid(
		[]validator.SimpleValidation{
			{
				Parameter: r.URL.Query().Get("page"),
				Validator: validator.StringPresent,
				ErrorMsg:  validator.MustProvidePageField,
			},
		}...)
}


func (m *MessageValidator) ValidateCreateBook(book *books.Book) error {
	return validator.FirstNonValid(
		[]validator.SimpleValidation{
			{
				Parameter: book.GetName(),
				Validator: validator.StringPresent,
				ErrorMsg:  validator.MustProvideNameField,
			},
			{
				Parameter: book.GetAuthor(),
				Validator: validator.StringPresent,
				ErrorMsg:  validator.MustProvideNameField,
			},
			{
				Parameter: book.GetPrice(),
				Validator: validator.PositiveInt,
				ErrorMsg:  validator.MustProvidePriceField,
			},
		}...)
}
