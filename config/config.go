package config

type Config struct {
	AppConfig   *AppConfig
	BooksConfig *BooksConfig
	LogConfig   *LogConfig
}

func DevConfig() *Config {
	return &Config{
		AppConfig:   devAppConfig(),
		BooksConfig: devBooksConfig(),
		LogConfig:   devLogConfig(),
	}
}
