package books

type BookSearch struct {
	Name string
}

func (s *BookSearch) GetName() string {
	if s == nil {
		return ""
	}
	return s.Name
}
