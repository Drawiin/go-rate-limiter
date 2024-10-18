package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeStorage struct {
	Storage map[string]map[string]int64
}

func (f *FakeStorage) Set(key string, timeWindow, requestCount int64) {
	if _, exists := f.Storage[key]; !exists {
		f.Storage[key] = make(map[string]int64)
	}
	f.Storage[key]["timeWindow"] = timeWindow
	f.Storage[key]["requestCount"] = requestCount
}

func (f *FakeStorage) Get(key string) (int64, int64) {
	values, exists := f.Storage[key]
	if !exists {
		return 0, 0
	}
	return values["timeWindow"], values["requestCount"]
}

func TestRateLimitService_ShouldThrottle(t *testing.T) {
	storage := make(map[string]map[string]int64)
	mockStorage := &FakeStorage{Storage: storage}
	config := RateLimitConfig{
		LimitPerIp:          5,
		LimitPerToken:       10,
		TimeWindowInSeconds: 60,
	}
	service := NewRateLimitService(mockStorage, config)

	tests := []struct {
		ip           string
		token        string
		time         int64
		expected     bool
		description  string
	}{
		{"192.168.1.1", "", 100, false, "First request should not throttle"},
		{"192.168.1.1", "", 100, false, "Second request within limit should not throttle"},
		{"192.168.1.1", "", 100, false, "Third request within limit should not throttle"},
		{"192.168.1.1", "", 100, false, "Fourth request within limit should not throttle"},
		{"192.168.1.1", "", 100, false, "Fifth request within limit should not throttle"},
		{"192.168.1.1", "", 100, true, "Sixth request should throttle"},
	}

	for _, test := range tests {
		result := service.ShouldThrottle(test.ip, test.token, test.time)
		assert.Equal(t, test.expected, result, test.description)
	}
}

func TestRateLimitService_ShouldThrottle_ByToken(t *testing.T) {
	storage := make(map[string]map[string]int64)
	mockStorage := &FakeStorage{Storage: storage}
	config := RateLimitConfig{
		LimitPerIp:          5,
		LimitPerToken:       3, // Set a lower limit for tokens to test throttling
		TimeWindowInSeconds: 60,
	}
	service := NewRateLimitService(mockStorage, config)

	tests := []struct {
		ip           string
		token        string
		time         int64
		expected     bool
		description  string
	}{
		{"192.168.1.1", "token123", 100, false, "First request with token should not throttle"},
		{"192.168.1.1", "token123", 100, false, "Second request with token should not throttle"},
		{"192.168.1.1", "token123", 100, false, "Third request with token should not throttle"},
		{"192.168.1.1", "token123", 100, true, "Fourth request with token should throttle"},
	}

	for _, test := range tests {
		result := service.ShouldThrottle(test.ip, test.token, test.time)
		assert.Equal(t, test.expected, result, test.description)
	}
}

func TestRateLimitService_ShouldNotThrottle_AfterTimeWindow(t *testing.T) {
	storage := make(map[string]map[string]int64)
	mockStorage := &FakeStorage{Storage: storage}
	config := RateLimitConfig{
		LimitPerIp:          5,
		LimitPerToken:       3,
		TimeWindowInSeconds: 60,
	}
	service := NewRateLimitService(mockStorage, config)

	tests := []struct {
		ip           string
		token        string
		time         int64
		expected     bool
		description  string
	}{
		{"192.168.1.1", "token123", 100, false, "First request with token should not throttle"},
		{"192.168.1.1", "token123", 100, false, "Second request with token should not throttle"},
		{"192.168.1.1", "token123", 100, false, "Third request with token should not throttle"},
		{"192.168.1.1", "token123", 100, true, "Fourth request with token should throttle"},
		{"192.168.1.1", "token123", 200, false, "Request after time window should not throttle"},
	}

	for _, test := range tests {
		result := service.ShouldThrottle(test.ip, test.token, test.time)
		assert.Equal(t, test.expected, result, test.description)
	}
}