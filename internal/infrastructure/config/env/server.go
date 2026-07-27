package env

import (
	"fmt"
	"net"

	"github.com/caarlos0/env/v11"
	errorapp "github.com/goNiki/subservice/internal/models/errorApp"
)

type serverEnvConfig struct {
	Host string `env:"SERVER_HOST,required"`
	Port string `env:"SERVER_PORT,required"`
}

type serverConfig struct {
	raw serverEnvConfig
}

func NewServerConfig() (*serverConfig, error) {
	var raw serverEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("%w: %v", errorapp.ErrParseServerConfig, err)
	}

	return &serverConfig{
		raw: raw,
	}, nil
}

func (cfg *serverConfig) Host() string {
	return cfg.raw.Host
}

func (cfg *serverConfig) Port() string {
	return cfg.raw.Port
}

func (cfg *serverConfig) Adress() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}
