package users

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/security"
	"bookstore/time"
)

type Backend interface {
	Create(user *User) (*User, error)
	Get(name, password string) (*User, error)
}

type Factory struct {
	Backend Backend
	time    *time.Factory
	log     *log.Factory
}

var factory = map[config.UsersBackend]func() Backend{
	config.UsersMemoryBackend: newMemoryBackend,
}

func NewUserFactory(config *config.UsersConfig, log *log.Factory, time *time.Factory) *Factory {
	return &Factory{
		time:    time,
		log:     log,
		Backend: factory[config.Backend](),
	}
}

func (u *Factory) CreateUser(user *User) (*User, error) {
	oldPdw := user.Password
	pwd, err := security.HashAndSalt(user.Password)
	if err != nil {
		u.log.Error("error hashing password", err)
	}
	user.Password = pwd
	user, err = u.Backend.Create(user)
	if err != nil {
		u.log.Error("error saving user", err)
	}
	user.Password = oldPdw
	return user, nil
}
