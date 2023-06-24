package wrapper

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vanyovan/mini-wallet.git/internal/helper"
)

type SqlWrapper struct {
	db *sql.DB
}

func NewSqlWrapper(db *sql.DB) SqlWrapper {
	return SqlWrapper{
		db: db,
	}
}

func (s *SqlWrapper) BeginTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, helper.TxKey, tx)
	err = fn(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func FromContext(ctx context.Context) (*sql.Tx, error) {
	currentTx, ok := ctx.Value(helper.TxKey).(*sql.Tx)
	if !ok {
		return nil, errors.New("failed to get tx from context")
	}
	if currentTx == nil {
		return nil, errors.New("failed to get tx from context")
	}
	return currentTx, nil
}
