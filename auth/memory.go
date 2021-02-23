package auth

import (
	"errors"
	"time"

	"bookstore/config"
	"bookstore/log"
	timeb "bookstore/time"
)

const sessionName = "bookstoreSession"

type memoryBackend struct {
	sessions map[string]*Session
	conf     *config.AuthConfig
	log      *log.Factory
	time     *timeb.Factory
}

func (m *memoryBackend) Delete(token string) error {
	delete(m.sessions, token)
	return nil
}

func (m *memoryBackend) Get(token string) (*Session, error) {
	session, found := m.sessions[token]
	if !found {
		return nil, errors.New("unauthenticated")
	}
	now := m.time.Now()

	if now.After(session.CreatedAt.Add(time.Minute * time.Duration(m.conf.SessionLenInMinutes))) {
		delete(m.sessions, token)
		return nil, errors.New("expired")
	}

	return session, nil
}

func (m *memoryBackend) Create(session *Session) (*Session, error) {
	m.sessions[session.Token] = session
	return session, nil
}

func newMemoryBackend(conf *config.AuthConfig, log *log.Factory, time *timeb.Factory) Backend {
	return &memoryBackend{
		sessions: map[string]*Session{},
		conf: conf,
		log:  log,
		time: time,
	}
}
