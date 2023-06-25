package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/repo/wrapper"
)

type WalletRepo interface {
	CreateWallet(ctx context.Context, userId string) (result entity.Wallet, err error)
	UpdateStatusByUserId(ctx context.Context, status string, userId string) (updatedAt time.Time, err error)
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
	tx, err := wrapper.FromContext(ctx)
	if tx == nil || err != nil {
		tx, err = r.db.Begin()
		if err != nil {
			tx.Rollback()
			return result, errors.New("failed to begin database transaction")
		}
	}
	guuid := helper.GenerateGuuid()
	timeNow := time.Now()
	_, err = tx.ExecContext(ctx, "INSERT INTO mst_wallet (wallet_id, owned_by, status, enabled_at, balance) VALUES (?, ?, ?, ?, ?)", guuid, userId, helper.ConstantEnabled, timeNow, helper.ConstantDefaultFloat64)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("failed to create wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to commit database transaction")
	}

	result = entity.Wallet{
		WalletId:  guuid,
		OwnedBy:   userId,
		Status:    helper.ConstantEnabled,
		EnabledAt: &timeNow,
		Balance:   helper.ConstantDefaultFloat64,
	}
	return result, nil
}

func (r *Repo) UpdateStatusByUserId(ctx context.Context, status string, userId string) (updatedAt time.Time, err error) {
	tx, err := wrapper.FromContext(ctx)
	if tx == nil || err != nil {
		tx, err = r.db.Begin()
		if err != nil {
			tx.Rollback()
			return time.Time{}, errors.New("failed to begin database transaction")
		}
	}

	timeNow := time.Now()
	if status == helper.ConstantEnabled {
		_, err = tx.Exec("UPDATE mst_wallet set status = ?, enabled_at = ? where owned_by = ?", status, timeNow, userId)
	} else {
		_, err = tx.Exec("UPDATE mst_wallet set status = ?, disabled_at = ? where owned_by = ?", status, timeNow, userId)
	}
	if err != nil {
		tx.Rollback()
		return time.Time{}, fmt.Errorf("failed to update wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return time.Time{}, errors.New("failed to commit database transaction")
	}

	return timeNow, nil
}
