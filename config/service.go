package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

func New() (*Service, error) {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	return &Service{
		envConfig: envConfig,
	}, nil
}

type Service struct {
	envConfig envConfig
}

func (s *Service) DBConnectionString() string {
	connFragments := []string{
		fmt.Sprintf("host=%s", s.envConfig.DBHostname),
		fmt.Sprintf("port=%d", s.envConfig.DBPort),
		fmt.Sprintf("dbname=%s", s.envConfig.DBDatabase),
	}

	if s.envConfig.DBSSLMode {
		connFragments = append(connFragments, "sslmode=require")
	} else {
		connFragments = append(connFragments, "sslmode=disable")
	}

	if s.envConfig.DBUsername != "" {
		connFragments = append(connFragments, fmt.Sprintf("user=%s", s.envConfig.DBUsername))
	}

	if s.envConfig.DBPassword != "" {
		connFragments = append(connFragments, fmt.Sprintf("password=%s", s.envConfig.DBPassword))
	}

	return strings.Join(connFragments, " ")
}

func (s *Service) SingleUser() bool {
	return s.envConfig.SingleUser
}

func (s *Service) JWTSecret() string {
	return s.envConfig.JWTSecret
}

func (s *Service) SessionSecret() string {
	return s.envConfig.SessionSecret
}

type envConfig struct {
	DBHostname    string `env:"DB_HOSTNAME,required"`
	DBPort        int64  `env:"DB_PORT,required"`
	DBUsername    string `env:"DB_USERNAME"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBDatabase    string `env:"DB_DATABASE,required"`
	DBSSLMode     bool   `env:"DB_SSL_MODE"`
	JWTSecret     string `env:"JWT_SECRET" envDefault:"please-change-me"`
	SessionSecret string `env:"SESSION_SECRET" envDefault:"change-me-please"`
	SingleUser    bool   `env:"SINGLE_USER"`
}
