package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	WebServerHost    string `mapstructure:"WEBSERVER_HOST"`
	WebServerPort    string `mapstructure:"WEBSERVER_PORT"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPort        string `mapstructure:"REDIS_PORT"`
	RateLimitByIp    string `mapstructure:"RATE_LIMIT_BY_IP"`
	RateLimitByToken string `mapstructure:"RATE_LIMIT_BY_TOKEN"`
	RateLimitWindow  string `mapstructure:"RATE_LIMIT_WINDOW"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	// Environment variables will override the .env values eg export DB_DRIVER=postgres will override the DB_DRIVER value in .env
	// enabling us to change some values without changing the .env file for testing purposes' viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
