package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckWebsite(t *testing.T) {
	service := &WebsiteService{}
	result := service.CheckWebsite("https://google.com")

	assert.NotEmpty(t, result.URL)
	assert.NotZero(t, result.CheckedAt)
	assert.Greater(t, result.Duration, time.Duration(0))
}
