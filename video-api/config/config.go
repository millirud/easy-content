package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App `yaml:"app"`
		Log
		RabbitMq
		Transform `yaml:"transform"`
		Storage   `yaml:"storage"`
	}

	// App -.
	App struct {
		Name        string `env-required:"true" yaml:"app_name"`
		Description string `env-required:"true" yaml:"app_description"`
		Version     string `env-required:"true" yaml:"app_version"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL"`
	}

	// MQ
	RabbitMq struct {
		User     string `env-required:"true" env:"RABBITMQ_USER"`
		Password string `env-required:"true" env:"RABBITMQ_PASS"`
		Host     string `env-required:"true" env:"RABBITMQ_HOST"`
		Port     int    `env-required:"true" env:"RABBITMQ_PORT"`
	}

	Transform struct {
		Queue string `env-required:"true" yaml:"transform_queue"`
	}

	Storage struct {
		Url     string `env-required:"true" env:"STORAGE_URL"`
		Timeout int    `env-required:"true" yaml:"storage_timeout"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
