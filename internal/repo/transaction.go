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

type TransactionRepo interface {
	UpdateWalletBalanceByUserId(ctx context.Context, amount float64, userId string) (result entity.Wallet, err error)
	CreateWalletTransaction(ctx context.Context, walletId string, status string, typeTransaction string, amount float64, referenceId string) (err error)
	GetTransactionByWalletId(walletId string) (result []entity.Transaction, err error)
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) UpdateWalletBalanceByUserId(ctx context.Context, amount float64, userId string) (result entity.Wallet, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to begin database transaction")
	}
	_, err = tx.ExecContext(ctx, "UPDATE mst_wallet set balance = ? where owned_by = ?", amount, userId)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("failed to update wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to commit database transaction")
	}

	return result, nil
}

func (r *Repo) CreateWalletTransaction(ctx context.Context, walletId string, status string, typeTransaction string, amount float64, referenceId string) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to begin database transaction")
	}

	guuid := helper.GenerateGuuid()
	_, err = tx.ExecContext(ctx, "INSERT INTO trx_wallet (wallet_id, transaction_id, status, transacted_at, type, amount, reference_id) VALUES (?, ?, ?, ?, ?, ?, ?)", walletId, guuid, status, time.Now(), typeTransaction, amount, referenceId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to commit database transaction")
	}
	return nil
}

func (r *Repo) GetTransactionByWalletId(walletId string) (result []entity.Transaction, err error) {
	rows, err := r.db.Query("SELECT wallet_id, transaction_id, status, transacted_at, type, amount, reference_id FROM trx_wallet WHERE wallet_id = ?", walletId)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		transaction := &entity.Transaction{}
		err = rows.Scan(&transaction.WalletId, &transaction.TransactionId, &transaction.Status, &transaction.TransactedAt, &transaction.Type, &transaction.Amount, &transaction.ReferenceId)
		if err != nil {
			return result, err
		}
		result = append(result, *transaction)
	}

	return result, nil
}
