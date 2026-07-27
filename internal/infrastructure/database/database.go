package database

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/subservice/internal/infrastructure/config"
	errorapp "github.com/goNiki/subservice/internal/models/errorApp"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func InitDatabase(cfg config.Postgres) (*DB, error) {
	poolcfg, err := pgxpool.ParseConfig(cfg.Adress())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorapp.ErrParseDatabaseConfig, err)
	}

	poolcfg.MaxConns = cfg.MaxConns()
	poolcfg.MinConns = cfg.MinConns()
	poolcfg.MaxConnIdleTime = cfg.MaxConnIdleTime()
	poolcfg.MaxConnLifetime = cfg.MaxConnLifeTime()
	poolcfg.HealthCheckPeriod = cfg.HealthCheckPerion()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolcfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errorapp.ErrCreateNewConnect, err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%w, %v", errorapp.ErrConnectDatabase, err)
	}

	return &DB{
		Pool: pool,
	}, nil
}

type QueryExecutor interface {
	QueryRow(ctx context.Context, sql string, arg ...any) pgx.Row
	Query(ctx context.Context, sql string, arg ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arg ...any) (pgconn.CommandTag, error)
}
