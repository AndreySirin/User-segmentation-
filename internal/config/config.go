package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig   ServerConfig   `yaml:"ServerConfig"`
	DatabaseConfig DatabaseConfig `yaml:"DatabaseConfig"`
	LoggerConfig   LoggerConfig   `yaml:"loggerConfig"`
}

type ServerConfig struct {
	Port            string        `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
	Module          string        `yaml:"Module"`
}

type DatabaseConfig struct {
	DbName   string `yaml:"DbName"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Address  string `yaml:"Address"`
}
type LoggerConfig struct {
	Debug bool `yaml:"debug"`
	Error bool `yaml:"error"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}
	return &config, nil
}

func PathConfig() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	exeDir := filepath.Dir(exePath)

	configPath := filepath.Join(exeDir, "config.yaml")

	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute config path: %w", err)
	}

	return absConfigPath, nil
}
