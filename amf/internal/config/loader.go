package type

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v6"
)

func LoadConfig() (Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return cfg, fmt.Errorf("unable to load .env file: %w", err)
	}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to parse environment variables: %w", err)
	}
	return cfg, nil
}