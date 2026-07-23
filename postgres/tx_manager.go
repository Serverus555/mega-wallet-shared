package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type TxManager struct {
	pool   *pgxpool.Pool
	logger zerolog.Logger
}

type TxExecutor interface {
	Exec(ctx context.Context, fn func(txCtx context.Context) error) error
}

type DbProvider interface {
	GetDB(ctx context.Context) DB
}

type txKey struct{}

type DB interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewTxManager(pool *pgxpool.Pool, logger zerolog.Logger) *TxManager {
	return &TxManager{
		pool:   pool,
		logger: logger,
	}
}

func (m *TxManager) GetDB(ctx context.Context) DB {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return m.pool
}

func (m *TxManager) Exec(ctx context.Context, fn func(txCtx context.Context) error) error {
	var tx pgx.Tx
	var err error

	if currentTx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		m.logger.Trace().Msg("TxManager: Begin nested")
		tx, err = currentTx.Begin(ctx)
	} else {
		m.logger.Trace().Msg("TxManager: Begin new")
		tx, err = m.pool.Begin(ctx)
	}

	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			m.logger.Trace().Msg("TxManager: Panic. Rollback")
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	if err = fn(context.WithValue(ctx, txKey{}, tx)); err != nil {
		m.logger.Trace().Msg("TxManager: Rollback")
		_ = tx.Rollback(ctx)
		return err
	}

	m.logger.Trace().Msg("TxManager: Commit")
	return tx.Commit(ctx)
}
