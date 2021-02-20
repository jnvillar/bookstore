package config

type Config struct {
	BooksConfig *BooksConfig
	LogConfig   *LogConfig
}

func DevConfig() *Config {
	return &Config{
		BooksConfig: devBooksConfig(),
		LogConfig:   devLogConfig(),
	}
}
