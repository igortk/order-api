package config

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type (
	Config struct {
		RabbitConfig RabbitConfig `envPrefix:"RMQ_"`
		LoggerConfig LoggerConfig `envPrefix:"LOG_"`
		HttpConfig   HttpConfig   `envPrefix:"HTTP_"`
	}

	RabbitConfig struct {
		Host              string `env:"HOST" envDefault:"localhost"`
		Port              int    `env:"PORT"  envDefault:"5672"`
		Username          string `env:"USERNAME" envDefault:"guest"`
		Password          string `env:"PASSWORD" envDefault:"guest"`
		VirtualHost       string `env:"VIRTUAL_HOST" envDefault:"/"`
		ReconnectAttempts uint   `env:"RECONNECT_ATTEMPTS" envDefault:"5"`
	}

	LoggerConfig struct {
		Level string `env:"LEVEL" envDefault:"INFO"`
	}

	HttpConfig struct {
		Port int `env:"PORT"  envDefault:"8081"`
	}
)

func GetConfig() (*Config, error) {
	log.Info(fmt.Sprintf("getting config..."))
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("can't parse config")
	}
	return &cfg, nil
}
