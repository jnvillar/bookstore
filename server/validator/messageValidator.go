package validator

import (
	"net/http"

	"bookstore/books"
	"bookstore/login"
)

// MessageValidator implements the Validator interface
type MessageValidator struct {
}

func NewMessageValidator() *MessageValidator {
	return &MessageValidator{}
}

func (m *MessageValidator) ValidateUpdateBook(params map[string]interface{}) error {
	return FirstNonValid(
		[]SimpleValidation{
			{
				Parameter: params["id"],
				Validator: StringPresent,
				ErrorMsg:  MustProvideIDField,
			},
		}...)
}

func (m *MessageValidator) ValidateListBooks(r *http.Request) error {
	return FirstNonValid(
		[]SimpleValidation{
			{
				Parameter: r.URL.Query().Get("page"),
				Validator: StringPresent,
				ErrorMsg:  MustProvidePageField,
			},
		}...)
}

func (m *MessageValidator) ValidateCreateBook(book *books.Book) error {
	return FirstNonValid(
		[]SimpleValidation{
			{
				Parameter: book.GetName(),
				Validator: StringPresent,
				ErrorMsg:  MustProvideNameField,
			},
			{
				Parameter: book.GetPrice(),
				Validator: PositiveInt,
				ErrorMsg:  MustProvidePriceField,
			},
		}...)
}

func (m *MessageValidator) ValidateLogin(login *login.Request) error {
	return FirstNonValid(
		[]SimpleValidation{
			{
				Parameter: login.GetUserName(),
				Validator: StringPresent,
				ErrorMsg:  MustProvideUserNameField,
			},
			{
				Parameter: login.GetPassword(),
				Validator: StringPresent,
				ErrorMsg:  MustProvidePasswordField,
			},
		}...)
}
