package users

import (
	"errors"
)

type memoryBackend struct {
	users []*User
}

func (m *memoryBackend) Create(user *User) (*User, error) {
	u := &User{
		ID:       user.ID,
		UserName: user.UserName,
		Password: user.Password,
		Role:     user.Role,
	}
	m.users = append(m.users, u)
	return user, nil
}

func (m *memoryBackend) Get(userName string) (*User, error) {
	for _, user := range m.users {
		if user.UserName == userName {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func newMemoryBackend() Backend {

	return &memoryBackend{
		users: []*User{},
	}
}
