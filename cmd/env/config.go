package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host   string `envconfig:"HOST" default:"localhost"`
	Port   string `envconfig:"PORT" default:"8080"`
	DBHost string `envconfig:"DB_HOST"`
	DBPort string `envconfig:"DB_PORT"`
	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASSWORD"`
	DBName string `envconfig:"DB_NAME"`
}

func LoadAppConfig() (AppConfig, error) {
	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		return AppConfig{}, fmt.Errorf("failed to parse configuration from environment: %w", err)
	}

	return config, nil
}
