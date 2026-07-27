package txmanager

import (
	"context"

	"github.com/goNiki/subservice/internal/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

type TxManager struct {
	db *database.DB
}

func NewTxManager(db *database.DB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (tm *TxManager) WithTX(
	ctx context.Context, opts pgx.TxOptions,
	fn func(ctx context.Context, q database.QueryExecutor) error,
) error {
	tx, err := tm.db.Pool.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = fn(ctx, tx)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
