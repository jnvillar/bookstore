package config

type Config struct {
	AppConfig      *AppConfig
	BooksConfig    *BooksConfig
	LogConfig      *LogConfig
	UsersConfig    *UsersConfig
	SessionsConfig *SessionsConfig
}

func DevConfig() *Config {
	return &Config{
		AppConfig:      devAppConfig(),
		BooksConfig:    devBooksConfig(),
		LogConfig:      devLogConfig(),
		UsersConfig:    devUsersConfig(),
		SessionsConfig: devSessionsConfig(),
	}
}
