package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

type Config struct {
	ListenerConfig listenerConfig `yaml:"listener"`
	Logger         logger         `yaml:"logger"`
}

func NewConfig(configPathEnv string) (*Config, error) {
	cfg, err := loadConfig(configPathEnv)
	if err != nil {
		return nil, errors.Wrap(err, "init load")
	}

	err = cfg.validateConfig()
	if err != nil {
		return nil, errors.Wrap(err, "init validate")
	}

	return cfg, nil
}

func loadConfig(configPathEnv string) (*Config, error) {
	cfg := &Config{}

	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		return nil, errors.New("error reading ENV file path")
	}
	data, err := os.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return nil, errors.Wrap(err, "error reading YAML file")
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {

		return nil, errors.Wrap(err, "error unmarshalling YAML")
	}

	return cfg, nil
}

func (cfg *Config) validateConfig() error {
	err := cfg.ListenerConfig.validate()
	if err != nil {
		return errors.Wrap(err, "validateConfig")
	}
	err = cfg.ListenerConfig.validate()
	if err != nil {
		return errors.Wrap(err, "validateConfig")
	}
	return nil
}
