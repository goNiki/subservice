package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	errorapp "github.com/goNiki/subservice/internal/models/errorApp"
)

type loggerEnvConfig struct {
	Level    string `env:"LOGGER_LEVEL,required"`
	Format   string `env:"LOGGER_FORMAT,required"`
	Output   string `env:"LOGGER_OUTPUT,required"`
	FilePath string `env:"LOGGER_FILE_PATH,required"`
}

type loggerConfig struct {
	raw loggerEnvConfig
}

func NewLoggerConfig() (*loggerConfig, error) {
	var raw loggerEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, fmt.Errorf("%w: %v", errorapp.ErrParseLoggerConfig, err)
	}

	return &loggerConfig{
		raw: raw,
	}, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.raw.Level
}

func (cfg *loggerConfig) Format() string {
	return cfg.raw.Format
}

func (cfg *loggerConfig) Output() string {
	return cfg.raw.Output
}

func (cfg *loggerConfig) FilePath() string {
	return cfg.raw.FilePath
}
