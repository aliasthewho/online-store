package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() *Config {
	var cfg Config
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error leyendo config: %v", err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error parseando config: %v", err)
	}
	return &cfg
}
