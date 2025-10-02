package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config 和 SourceConfig 结构体需要更新
type Config struct {
	Port    int            `yaml:"port"`
	Sources []SourceConfig `yaml:"sources"`
}

type SourceConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

func LoadConfig(cfg string) *Config {
	yamlFile, err := os.ReadFile(cfg)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	return &config
}
