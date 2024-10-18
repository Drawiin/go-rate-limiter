package services

import (
	"fmt"
	"go-rate-limiter/internal/infra"
)

type RateLimitService struct {
	storage infra.KeyValueStore
	config  RateLimitConfig
}

type RateLimitConfig struct {
	LimitPerIp          int
	LimitPerToken       int
	TimeWindowInSeconds int
}

func NewRateLimitService(storage infra.KeyValueStore, config RateLimitConfig) *RateLimitService {
	return &RateLimitService{
		storage: storage,
		config:  config,
	}
}

func (s *RateLimitService) ShouldThrottle(ip, token string, time int64) bool {
	key := s.getKey(ip, token)
	limit := s.getLimit(ip, token)
	currentWindow := time / int64(s.config.TimeWindowInSeconds)
	timeWindow, requestCount := s.storage.Get(key)
	fmt.Printf("Key: %s, Limit: %d, Current Window: %d, Time Window: %d, Request Count: %d\n", key, limit, currentWindow, timeWindow, requestCount)
	if timeWindow == 0 {
		fmt.Printf("First request for key %v\n", key)
		s.storage.Set(key, currentWindow, 1)
		return false
	}

	// If the windows are different the windows has expired so we need to reset the request count
	if timeWindow != currentWindow {
		fmt.Printf("Window expired reseting count\n")
		s.storage.Set(key, currentWindow, 1)
		return false
	}

	if requestCount >= limit {
		return true
	}

	s.storage.Set(key, currentWindow, requestCount+1)

	return false
}

func (s *RateLimitService) getKey(ip, token string) string {
	return fmt.Sprintf("%s:%s", ip, token)
}

func (s *RateLimitService) getLimit(ip, token string) int64 {
	if token != "" {
		return int64(s.config.LimitPerToken)
	}

	return int64(s.config.LimitPerIp)
}