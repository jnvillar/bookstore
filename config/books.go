package config

type Backend = int

const (
	MemoryBackend Backend = iota
	PostgresBackend
)

type BooksConfig struct {
	Backend Backend
}

func devBooksConfig() *BooksConfig {
	return &BooksConfig{
		Backend: MemoryBackend,
	}
}
