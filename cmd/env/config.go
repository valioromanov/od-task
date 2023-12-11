package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	Host   string `envconfig:"HOST" default:"localhost"`
	Port   int    `envconfig:"PORT" default:"8080"`
	DBHost string `envconfig:"DB_HOST"`
	DBPort int    `envconfig:"DB_PORT"`
	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASSWORD"`
	DBName string `envconfig:"DB_NAME"`
}

func LoadAppConfig() (AppConfig, error) {
	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		logrus.Error("error while binding to 'AppConfig': ", err)
		return AppConfig{}, fmt.Errorf("failed to parse configuration from environment: %w", err)
	}

	return config, nil
}
