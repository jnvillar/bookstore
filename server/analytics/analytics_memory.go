package analytics

import "bookstore/config"

type memoryBackend struct {
	visits map[string]bool
}

func (m *memoryBackend) AddVisit(visit *Visit) error {
	m.visits[visit.Ip] = true
	return nil
}

func (m *memoryBackend) GetAnalytics() (*Info, error) {
	return &Info{Visits: len(m.visits)}, nil
}

func newMemoryBackend(conf *config.AnalyticsConfig) Backend {
	return &memoryBackend{visits: map[string]bool{}}
}
