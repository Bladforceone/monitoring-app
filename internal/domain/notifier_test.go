package domain

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) SendAlert(siteURL string, statusCode int, err error) {
	m.Called(siteURL, statusCode, err)
}

func TestSendAlert(t *testing.T) {
	mockNotifier := new(MockNotifier)
	mockNotifier.On("SendAlert", "https://example.com", 500, nil).Return()

	mockNotifier.SendAlert("https://example.com", 500, nil)
	mockNotifier.AssertCalled(t, "SendAlert", "https://example.com", 500, nil)
}
