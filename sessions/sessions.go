package sessions

import (
	"bookstore/config"
)

type Backend interface {
	Get(token string) (*Session, error)
	Store(token string) error
}

type Factory struct {
	Backend Backend
}

var factory = map[config.UsersBackend]func() Backend{
	config.UsersMemoryBackend: newMemoryBackend,
}

func NewSessionsFactory(conf config.UsersConfig) *Factory {

}
