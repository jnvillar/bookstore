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
	MustProvidePageField  = mustProvideFieldErrorMessage("Page")
	MustProvideNameField  = mustProvideFieldErrorMessage("Name")
	MustProvidePriceField = mustProvidePositiveIntErrorMessage("Price")
)
