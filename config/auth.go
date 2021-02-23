package config

type AuthBackend = int

const (
	AuthMemoryBackend AuthBackend = iota
)

type AuthConfig struct {
	Backend             AuthBackend
	SessionLenInMinutes int
	DisableAuth         bool
}

func devAuthConfig() *AuthConfig {
	return &AuthConfig{
		Backend:             AuthMemoryBackend,
		SessionLenInMinutes: 60,
		DisableAuth:         true,
	}
}
