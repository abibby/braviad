package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var cfg *Config

func Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	c := &Config{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return err
	}
	cfg = c
	return nil
}

func Get(v any) error {
	if cfg == nil {
		return errors.New("config not loaded")
	}
	return json.Unmarshal(cfg.Data, v)
}
