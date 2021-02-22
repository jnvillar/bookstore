package users

import (
	"errors"

	"bookstore/config"
	"bookstore/log"
	"bookstore/roles"
	"bookstore/security"
	"bookstore/time"

	"github.com/google/uuid"
)

type Backend interface {
	Create(user *User) (*User, error)
	Get(userName string) (*User, error)
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
	factory := &Factory{
		time:    time,
		log:     log,
		Backend: factory[config.Backend](),
	}
	factory.addAdminUser()
	return factory
}

func (u *Factory) addAdminUser() {
	_, err := u.Create(&User{
		ID:       uuid.New().String(),
		UserName: "admin",
		Password: "admin",
		Role:     roles.RoleAdmin,
	})
	if err != nil {
		u.log.Error("error creating admin user", err)
	}

}

func (u *Factory) Create(user *User) (*User, error) {
	oldPdw := user.Password
	pwd, err := security.HashAndSalt(user.Password)
	if err != nil {
		u.log.Error("error hashing password", err)
	}
	user.Password = pwd
	user, err = u.Backend.Create(user)
	if err != nil {
		u.log.Error("error creating user", err)
	}
	user.Password = oldPdw
	return user, nil
}

func (u *Factory) Get(username, password string) (*User, error) {
	user, err := u.Backend.Get(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if correctPassword := security.ComparePassWords(password, user.Password); !correctPassword {
		return nil, errors.New("incorrect password")
	}
	return user, nil
}
