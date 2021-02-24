package config

type BooksBackend = int

const (
	BooksMemoryBackend BooksBackend = iota
	BooksPostgresBackend
)

type BooksConfig struct {
	Backend  BooksBackend
	PageSize int
}

func devBooksConfig() *BooksConfig {
	return &BooksConfig{
		Backend:  BooksMemoryBackend,
		PageSize: 100,
	}
}
