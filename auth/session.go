package auth

import (
	"time"

	"bookstore/roles"
)

type Session struct {
	Token     string
	Role      roles.Role
	CreatedAt time.Time
}
