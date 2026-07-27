package config

import (
	"fmt"

	"github.com/goNiki/subservice/internal/infrastructure/config/env"
	errorapp "github.com/goNiki/subservice/internal/models/errorApp"
	"github.com/joho/godotenv"
)

type config struct {
	Server   Server
	Logger   Logger
	Postgres Postgres
}

func Load(path string) (*config, error) {

	if err := godotenv.Load(path); err != nil {
		return nil, fmt.Errorf("%w: %v", errorapp.ErrLoadEnv, err)
	}

	server, err := env.NewServerConfig()
	if err != nil {
		return nil, err
	}

	logger, err := env.NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return nil, err
	}

	return &config{
		Server:   server,
		Logger:   logger,
		Postgres: postgres,
	}, nil
}
