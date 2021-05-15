package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	AppPort string `yaml:"app_port"`
	Db struct {
		Postgres struct {
			Host string `yaml:"psql_host"`
			Port string `yaml:"psql_port"`
			User string `yaml:"psql_user"`
			Password string `yaml:"psql_password"`
			DbName string `yaml:"psql_db_name"`
			SslMode string `yaml:"psql_ssl_mode"`
		}
	}
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
