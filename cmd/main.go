package main

import (
	"log"
	"monitoring-app/config"
	"monitoring-app/internal/api"
	"monitoring-app/internal/domain"
	"monitoring-app/internal/repository"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".env", "config.yaml")
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	db, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	notifier := domain.NewTelegramNotifier(cfg.TelegramBotToken, cfg.TelegramChatID)

	service := domain.NewWebsiteService(notifier, cfg.Monitoring.Websites, cfg.Monitoring.Interval)

	log.Println("🔄 Запуск мониторинга сайтов...")
	go service.StartMonitoring()

	repo := repository.NewRepository(db)
	apiHandler := &api.APIHandler{
		Service: service,
		Repo:    repo,
	}

	mux := api.SetupRoutes(apiHandler)

	log.Println("✅ Автоматический мониторинг + API запущены на порту", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, mux))
}
