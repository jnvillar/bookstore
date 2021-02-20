package log

import "bookstore/config"

type Backend interface {
	Debug(log string)
	Info(log string)
}

type Factory struct {
	Backend
}

var factory = map[config.LogBackend]func(backend *config.LogConfig) Backend{
	config.LogStdOutBackend: newStdoutLogger,
}

func NewLoggerFactory(conf *config.LogConfig) *Factory {
	return &Factory{Backend: factory[conf.Backend](conf)}
}
