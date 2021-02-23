package config

type Config struct {
	AppConfig   *AppConfig
	BooksConfig *BooksConfig
	LogConfig   *LogConfig
	UsersConfig *UsersConfig
	AuthConfig  *AuthConfig
}

func DevConfig() *Config {
	return &Config{
		AppConfig:   devAppConfig(),
		BooksConfig: devBooksConfig(),
		LogConfig:   devLogConfig(),
		UsersConfig: devUsersConfig(),
		AuthConfig:  devAuthConfig(),
	}
}
