package config

import (
	"time"

	"github.com/goNiki/subservice/internal/infrastructure/config/env"
)

type Logger interface {
	Level() string
	Console() env.ConsoleChannel
	File() env.FileChannel
}

type Postgres interface {
	Adress() string
	Host() string
	Port() string
	User() string
	Password() string
	Name() string
	SslMode() string
	MaxConns() int32
	MinConns() int32
	MaxConnLifeTime() time.Duration
	MaxConnIdleTime() time.Duration
	HealthCheckPerion() time.Duration
}

type Server interface {
	Host() string
	Port() string
	Adress() string
}
