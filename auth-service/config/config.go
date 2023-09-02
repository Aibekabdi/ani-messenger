package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		API `yaml:"api"`
		PG  `yaml:"pg"`
		Log `yaml:"log"`
	}
	API struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	PG struct {
		URL string `env-required:"true" env:"PG_URL"`
	}
	Log struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
		Output string `yaml:"output"`
	}
)

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

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
