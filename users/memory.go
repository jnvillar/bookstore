package users

import (
	"errors"
)

type memoryBackend struct {
	users []*User
}

func (m *memoryBackend) Create(user *User) (*User, error) {
	m.users = append(m.users, user)
	return user, nil
}

func (m *memoryBackend) Get(name, password string) (*User, error) {
	for _, user := range m.users {
		if user.Name == name && user.Password == password {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func newMemoryBackend() Backend {

	return &memoryBackend{
		users: []*User{
			{Name: "admin", Password: "admin", Role: UserRoleAdmin},
		},
	}
}
