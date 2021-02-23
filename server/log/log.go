package log

import (
	"bookstore/config"
	"bookstore/time"
)

type Backend interface {
	Debug(log string)
	Info(log string)
	Error(log string, err error)
}

type Factory struct {
	Backend
	time *time.Factory
}

var factory = map[config.LogBackend]func(backend *config.LogConfig, time *time.Factory) Backend{
	config.LogStdOutBackend: newStdoutLogger,
}

func NewLoggerFactory(conf *config.LogConfig, time *time.Factory) *Factory {
	return &Factory{
		Backend: factory[conf.Backend](conf, time),
		time:    time,
	}
}
