package domain

import (
	"net/http"
	"time"
)

type WebsiteService struct {
	Notifier Notifier
	Websites []string
	Interval time.Duration
}

type Website struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Error      error
	CheckedAt  time.Time
}

func NewWebsiteService(notifier Notifier, websites []string, interval time.Duration) *WebsiteService {
	return &WebsiteService{
		Notifier: notifier,
		Websites: websites,
		Interval: interval,
	}
}

func (s *WebsiteService) StartMonitoring() {
	for {
		for _, site := range s.Websites {
			result := s.CheckWebsite(site)
			if result.Error != nil || result.StatusCode >= 500 {
				s.Notifier.SendAlert(result.URL, result.StatusCode, result.Error)
			}
		}
		time.Sleep(s.Interval)
	}
}

func (s *WebsiteService) CheckWebsite(url string) Website {
	start := time.Now()
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(url)
	duration := time.Since(start)

	website := Website{
		URL:       url,
		Duration:  duration,
		CheckedAt: time.Now(),
	}

	if err != nil {
		website.Error = err
		return website
	}
	defer resp.Body.Close()

	website.StatusCode = resp.StatusCode
	return website
}
