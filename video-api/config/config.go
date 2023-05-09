package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App `yaml:"app"`
		HTTP
		Log
	}

	// App -.
	App struct {
		Name        string `env-required:"true" yaml:"app_name"`
		Description string `env-required:"true" yaml:"app_description"`
		Version     string `env-required:"true" yaml:"app_version"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true"  env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL"`
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
