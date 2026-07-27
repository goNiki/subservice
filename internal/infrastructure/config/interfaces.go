package config

import "time"

type Logger interface {
	Level() string
	Format() string
	Output() string
	FilePath() string
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
