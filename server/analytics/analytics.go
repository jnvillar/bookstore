package analytics

import (
	"bookstore/config"
	"bookstore/log"
	"bookstore/time"
)

type Backend interface {
	AddVisit(visit *Visit) error
	GetAnalytics() (*Info, error)
}

type Factory struct {
	backend Backend
	log     *log.Factory
	time    *time.Factory
}

var factory = map[config.AnalyticsBackend]func(backend *config.AnalyticsConfig) Backend{
	config.AnalyticsMemoryBackend: newMemoryBackend,
}

func NewAnalyticsFactory(conf *config.AnalyticsConfig, log *log.Factory, time *time.Factory) *Factory {
	return &Factory{
		backend: factory[conf.Backend](conf),
		log:     log,
		time:    time,
	}
}

func (f *Factory) AddVisit(ip string) error {
	return f.backend.AddVisit(&Visit{Ip: ip})
}

func (f *Factory) GetAnalytics() (*Info, error) {
	return f.backend.GetAnalytics()
}
