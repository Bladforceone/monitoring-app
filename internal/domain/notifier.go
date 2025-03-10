package domain

import (
	"fmt"
	"net/http"
	"net/url"
)

type Notifier interface {
	SendAlert(siteURL string, statusCode int, err error)
}

type TelegramNotifier struct {
	BotToken string
	ChatID   string
}

func NewTelegramNotifier(botToken, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		BotToken: botToken,
		ChatID:   chatID,
	}
}

func (t *TelegramNotifier) SendAlert(siteURL string, statusCode int, err error) {
	message := fmt.Sprintf("⚠️ Проблема с сайтом: %s", siteURL)
	if err != nil {
		message += fmt.Sprintf("\nОшибка: %v", err)
	} else {
		message += fmt.Sprintf("\nСервер вернул статус-код: %d", statusCode)
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)
	data := url.Values{}
	data.Set("chat_id", t.ChatID)
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		fmt.Println("❌ Ошибка отправки в Telegram:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("✅ Уведомление отправлено в Telegram:", siteURL)
}
