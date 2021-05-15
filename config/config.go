package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
}

func (c *Config) Init(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
