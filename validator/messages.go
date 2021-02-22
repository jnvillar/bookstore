package validator

func errorMessage(reason, message string) string {
	return reason + ": " + message
}

func mustProvideFieldErrorMessage(field string) string {
	return errorMessage("Must provide field", field)
}

func mustProvidePositiveIntErrorMessage(field string) string {
	return errorMessage("Should be greater than zero", field)
}

var (
	MustProvideIDField       = mustProvideFieldErrorMessage("id")
	MustProvidePageField     = mustProvideFieldErrorMessage("page")
	MustProvideNameField     = mustProvideFieldErrorMessage("name")
	MustProvideUserNameField = mustProvideFieldErrorMessage("userName")
	MustProvidePasswordField = mustProvideFieldErrorMessage("password")
	MustProvidePriceField    = mustProvidePositiveIntErrorMessage("price")
)
