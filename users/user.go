package users

import "bookstore/roles"

type User struct {
	ID       string
	Name     string
	Password string
	Role     roles.Role
}
