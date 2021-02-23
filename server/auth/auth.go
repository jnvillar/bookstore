package auth

import (
	"errors"

	"bookstore/config"
	"bookstore/log"
	"bookstore/permissions"
	"bookstore/time"
	"bookstore/users"

	"github.com/google/uuid"
)

const AuthHeader = "auth"

type Backend interface {
	Get(token string) (*Session, error)
	Create(session *Session) (*Session, error)
	Delete(token string) error
}

type Factory struct {
	Backend Backend
	log     *log.Factory
	time    *time.Factory
	config  *config.AuthConfig
}

var factory = map[config.AuthBackend]func(
	authConfig *config.AuthConfig,
	log *log.Factory,
	time *time.Factory) Backend{
	config.AuthMemoryBackend: newMemoryBackend,
}

func NewSessionsFactory(conf *config.AuthConfig, log *log.Factory, time *time.Factory) *Factory {
	return &Factory{
		Backend: factory[conf.Backend](conf, log, time),
		log:     log,
		time:    time,
		config:  conf,
	}
}

func (f Factory) Validate(token string, action permissions.Action) error {
	if f.config.DisableAuth {
		return nil
	}
	session, err := f.Backend.Get(token)
	if err != nil {
		return err
	}
	if hasPermissions := permissions.HasPermissions(session.Role, action); !hasPermissions {
		return errors.New("unauthorized")
	}
	return nil
}

func (f *Factory) Create(user *users.User) (*Session, error) {
	session := &Session{
		Token:     uuid.New().String(),
		Role:      user.Role,
		CreatedAt: f.time.Now(),
	}
	return f.Backend.Create(session)
}

func (f *Factory) Delete(token string) error {
	return f.Backend.Delete(token)
}
