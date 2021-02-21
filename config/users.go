package config

type UsersBackend = int

const (
	UsersMemoryBackend UsersBackend = iota
	UsersPostgresBackend
)

type UsersConfig struct {
	Backend UsersBackend
}

func devUsersConfig() *UsersConfig {
	return &UsersConfig{
		Backend: UsersMemoryBackend,
	}
}
