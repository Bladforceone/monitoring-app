package config

import (
	"time"
)

type Config struct {
	DBUser     string `yaml:"-"` // Загружается из .env
	DBPassword string `yaml:"-"` // Загружается из .env
	DBHost     string `yaml:"-"` // Загружается из .env
	DBPort     string `yaml:"-"` // Загружается из .env
	DBName     string `yaml:"-"` // Загружается из .env
	DBSSLMode  string `yaml:"-"` // Загружается из .env
	ServerPort string `yaml:"-"` // Загружается из .env

	TelegramBotToken string `yaml:"-"` // Загружается из .env
	TelegramChatID   string `yaml:"-"` // Загружается из .env

	Monitoring struct {
		Interval time.Duration `yaml:"interval"`
		Websites []string      `yaml:"websites"`
	} `yaml:"monitoring"`
}
