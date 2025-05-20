package config

import (
	"os"
)

type Config struct {
	DSN string
}

// LoadConfig — просто получаем DSN из переменных окружения
func LoadConfig() *Config {
	return &Config{
		DSN: os.Getenv("DSN"),
	}
}

func (c *Config) GetDSN() string {
	return c.DSN
}
