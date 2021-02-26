package books

type BookSearch struct {
	Name       string
	Page       int
	PriceOrder PriceOrder
}

type PriceOrder = string

const (
	DESC PriceOrder = "desc"
	ASC  PriceOrder = "asc"
	NONE PriceOrder = ""
)

func StringToPriceOrder(s string) PriceOrder {
	switch s {
	case ASC:
		return ASC
	case DESC:
		return DESC
	default:
		return NONE
	}
}

func (s *BookSearch) GetName() string {
	if s == nil {
		return ""
	}
	return s.Name
}
