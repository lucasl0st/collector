package main

import "github.com/caarlos0/env/v9"

type Config struct {
	ApiKeys     string `env:"API_KEYS"`
	Port        uint16 `env:"PORT" envDefault:"80"`
	StoragePath string `env:"STORAGE_PATH" envDefault:"storage/"`
}

func ParseConfig() (*Config, error) {
	c := &Config{}
	err := env.Parse(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
