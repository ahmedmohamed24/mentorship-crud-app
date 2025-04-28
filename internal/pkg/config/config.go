package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", "10s")
	viper.SetDefault("server.write_timeout", "10s")
	viper.SetDefault("database.dsn", "")

	if path == "" {
		_, d, _, ok := runtime.Caller(0)
		if !ok {
			return nil, fmt.Errorf("failed to get caller info")
		}
		path = filepath.Join(filepath.Dir(d), "config.yaml")
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	envBindings := map[string]string{
		"server.port":          "SERVER_PORT",
		"server.read_timeout":  "SERVER_READ_TIMEOUT",
		"server.write_timeout": "SERVER_WRITE_TIMEOUT",
		"database.dsn":         "DATABASE_DSN",
	}

	for k, v := range envBindings {
		if err := viper.BindEnv(k, v); err != nil {
			return nil, fmt.Errorf("failed to bind env var %s: %w", v, err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate configuration
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Server.Port <= 0 || cfg.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", cfg.Server.Port)
	}

	if cfg.Server.ReadTimeout <= 0 {
		return fmt.Errorf("read timeout must be positive")
	}

	if cfg.Server.WriteTimeout <= 0 {
		return fmt.Errorf("write timeout must be positive")
	}

	if cfg.Database.DSN == "" {
		return fmt.Errorf("database DSN is required")
	}

	return nil
}
