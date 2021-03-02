package config

type AnalyticsBackend = int

const (
	AnalyticsMemoryBackend AnalyticsBackend = iota
)

type AnalyticsConfig struct {
	Backend AnalyticsBackend
}

func devAnalyticsConfig() *AnalyticsConfig {
	return &AnalyticsConfig{
		Backend: AnalyticsMemoryBackend,
	}
}
