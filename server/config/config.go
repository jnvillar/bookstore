package config

type Config struct {
	AppConfig       *AppConfig
	BooksConfig     *BooksConfig
	LogConfig       *LogConfig
	UsersConfig     *UsersConfig
	AuthConfig      *AuthConfig
	AnalyticsConfig *AnalyticsConfig
}

func DevConfig() *Config {
	return &Config{
		AnalyticsConfig: devAnalyticsConfig(),
		AppConfig:       devAppConfig(),
		BooksConfig:     devBooksConfig(),
		LogConfig:       devLogConfig(),
		UsersConfig:     devUsersConfig(),
		AuthConfig:      devAuthConfig(),
	}
}
