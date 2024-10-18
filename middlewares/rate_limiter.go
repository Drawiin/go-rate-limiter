package middlewares

import (
	"fmt"
	"go-rate-limiter/services"
	"net"
	"net/http"
	"time"
)

type RateLimiter struct {
	service *services.RateLimitService
}

func NewRateLimiter(service *services.RateLimitService) *RateLimiter {
	return &RateLimiter{service: service}
}

func (l RateLimiter) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		token := r.Header.Get("API_KEY")
		fmt.Printf("Rate limiting request ip %v token = %v\n:", ip, token)
		if l.service.ShouldThrottle(ip, token, time.Now().Unix()) {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}
