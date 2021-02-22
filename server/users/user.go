package users

import (
	"bookstore/roles"
)

type User struct {
	ID       string
	UserName string
	Password string
	Role     roles.Role
}
