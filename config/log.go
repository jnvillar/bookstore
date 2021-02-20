package config

type LogBackend = int

const (
	LogStdOutBackend LogBackend = iota
)

type LogConfig struct {
	Backend  LogBackend
	LogLevel int
}

func devLogConfig() *LogConfig {
	return &LogConfig{
		Backend:  LogStdOutBackend,
		LogLevel: 0,
	}
}
