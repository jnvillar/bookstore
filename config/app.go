package config

type AppConfig struct {
	Port int
}

func devAppConfig() *AppConfig {
	return &AppConfig{
		Port: 8080,
	}
}