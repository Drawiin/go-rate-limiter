package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-rate-limiter/config"
	"go-rate-limiter/controller"
	"go-rate-limiter/internal/infra"
	"go-rate-limiter/middlewares"
	"go-rate-limiter/services"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	appConfig := buildAppConfig()
	redisClient := buildRedisClient(appConfig)
	rateLimitService := buildRateLimitService(appConfig, redisClient)
	rateLimitMiddleware := middlewares.NewRateLimiter(rateLimitService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(rateLimitMiddleware.RateLimit)
	urlController := controller.NewUrlController(&map[string]string{}, appConfig.WebServerHost, appConfig.WebServerPort)

	r.Post("/url", urlController.CreateUrl)
	r.Get("/{urlId}", urlController.GetUrl)
	r.Get("/{urlId}/unwrap", urlController.GetUrlUnwrapped)

	http.ListenAndServe(fmt.Sprintf(":%s", appConfig.WebServerPort), r)
}

func buildRateLimitService(appConfig *config.Config, redisClient *redis.Client) *services.RateLimitService {
	rateByIp, err := strconv.ParseInt(appConfig.RateLimitByIp, 10, 32)
	rateByToken, err := strconv.ParseInt(appConfig.RateLimitByToken, 10, 32)
	timeWindow, err := strconv.ParseInt(appConfig.RateLimitWindow, 10, 32)
	if err != nil {
		panic(err)
	}

	store := infra.NewRedisStore(redisClient)

	return services.NewRateLimitService(store, services.RateLimitConfig{
		LimitPerIp:          int(rateByIp),
		LimitPerToken:       int(rateByToken),
		TimeWindowInSeconds: int(timeWindow),
	})
}

func buildRedisClient(appConfig *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%s", appConfig.RedisHost, appConfig.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func buildAppConfig() *config.Config {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading loadConfig: %v", err)
	}
	return loadConfig
}
