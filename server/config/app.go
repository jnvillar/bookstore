package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Port int
}

func getPort() int {
	port := os.Getenv("port")
	if port == "" {
		return 8080
	}
	p, _ := strconv.Atoi(port)
	return p
}

func devAppConfig() *AppConfig {
	return &AppConfig{
		Port: getPort(),
	}
}
