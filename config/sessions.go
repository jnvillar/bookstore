package config

type SessionsBackend = int

const (
	SessionsMemoryBackend SessionsBackend = iota
)

type SessionsConfig struct {
	Backend             SessionsBackend
	SessionsKey         string
	SessionLenInMinutes int
}

func devSessionsConfig() *SessionsConfig {
	return &SessionsConfig{
		Backend:             SessionsMemoryBackend,
		SessionsKey:         "test",
		SessionLenInMinutes: 60,
	}
}
