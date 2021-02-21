package sessions

import (
	"errors"
	"time"

	"bookstore/config"

	"github.com/gorilla/sessions"
)

const sessionName = "bookstoreSession"

type memoryBackend struct {
	sessions map[string]*Session
}

func (g *gorillaBackend) Get(token string) error {
	time, found := g.sessions[token]
	if !found {
		return errors.New("unathenticated")
	}
}

func (g gorillaBackend) Store() {
	panic("implement me")
}

func newGorillaBackend(conf *config.SessionsConfig) Backend {
	store := sessions.NewCookieStore([]byte(conf.SessionsKey))
	store.Options = &sessions.Options{
		MaxAge:   60 * conf.SessionLenInMinutes,
		HttpOnly: true,
	}
	return &gorillaBackend{
		store: store,
	}
}
