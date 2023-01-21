package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port             string `envconfig:"APP_PORT" validate:"required,startswith=:"`
	ServiceName      string `envconfig:"SERVICE_NAME" validate:"required"`
	StartAgentsCount int    `envconfig:"START_AGENTS_COUNT" validate:"required"`
}

func Read() (*Config, error) {
	_ = godotenv.Overload("./cmd/server/.env", "./cmd/server/.env.local")

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
