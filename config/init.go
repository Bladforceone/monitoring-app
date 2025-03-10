package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(envPath, yamlPath string) (*Config, error) {
	config := &Config{}

	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️ Файл .env не найден, используются системные переменные")
	}

	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
	config.DBSSLMode = os.Getenv("DB_SSLMODE")
	config.ServerPort = os.Getenv("SERVER_PORT")

	config.TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	config.TelegramChatID = os.Getenv("TELEGRAM_CHAT_ID")

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать config.yaml: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	log.Println("✅ Конфигурация загружена из .env и config.yaml")
	return config, nil
}
