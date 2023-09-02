package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
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
		URL string `env-required:"true"                 env:"PG_URL"`
	}
	Log struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
		Output string `yaml:"output"`
	}
)

func LoadConfig() (*Config, error) {
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
