package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
)

type WalletRepo interface {
	CreateWallet(ctx context.Context, userId string) (result entity.Wallet, err error)
	GetWalletByUserId(ctx context.Context, userId string) (result entity.Wallet, err error)
}

func NewWalletRepo(db *sql.DB) WalletRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetWalletByUserId(ctx context.Context, userId string) (result entity.Wallet, err error) {
	query := "SELECT wallet_id, owned_by, status, enabled_at, disabled_at, balance FROM mst_wallet WHERE owned_by = ?"
	row := r.db.QueryRow(query, userId)
	result = entity.Wallet{}
	err = row.Scan(&result.WalletId, &result.OwnedBy, &result.Status, &result.EnabledAt, &result.DisabledAt, &result.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		} else {
			fmt.Println("Failed to retrieve row:", err)
		}
		return result, err
	}
	return result, nil
}

func (r *Repo) CreateWallet(ctx context.Context, userId string) (result entity.Wallet, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to begin database transaction")
	}
	guuid := helper.GenerateGuuid()

	_, err = tx.ExecContext(ctx, "INSERT INTO mst_wallet (wallet_id, owned_by, status, enabled_at, balance) VALUES (?, ?, ?, ?, ?)", guuid, userId, helper.ConstantEnabled, time.Now(), helper.ConstantDefaultInt)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("failed to create wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to commit database transaction")
	}

	return result, nil
}
