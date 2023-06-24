package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
)

type TransactionService struct {
	UserRepo        repo.UserRepo
	WalletRepo      repo.WalletRepo
	TransactionRepo repo.TransactionRepo
	Db              *sql.DB
}

type TransactionServiceProvider interface {
	CreateDepositWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error)
	ViewWalletTransaction(ctx context.Context, user entity.User) (result entity.Transaction, err error)
}

func NewTransactionService(UserRepo repo.UserRepo, WalletRepo repo.WalletRepo, TransactionRepo repo.TransactionRepo, Db *sql.DB) TransactionService {
	return TransactionService{
		UserRepo:        UserRepo,
		WalletRepo:      WalletRepo,
		TransactionRepo: TransactionRepo,
		Db:              Db,
	}
}

func (uc *TransactionService) CreateDepositWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error) {
	// get wallet.
	wallet, err := uc.WalletRepo.GetWalletByUserId(ctx, user.CustomerXid)
	if helper.IsStructEmpty(wallet) || wallet.Status == helper.ConstantDisabled {
		return result, errors.New("wallet not found or wallet is disabled")
	}
	if err != nil {
		return result, err
	}

	var status = helper.ConstantSuccess

	//count balance
	wallet.Balance = wallet.Balance + param.Amount

	//update wallet amount
	updateWallet, err := uc.TransactionRepo.UpdateWalletBalanceByUserId(ctx, wallet.Balance, user.CustomerXid)
	fmt.Println(updateWallet)
	if err != nil {
		status = helper.ConstantFailed
		return result, err
	}

	//create transaction deposit wallet
	err = uc.TransactionRepo.CreateWalletTransaction(ctx, wallet.WalletId, status, helper.ConstantDeposit, param.Amount, param.ReferenceId)
	return wallet, err
}

func (uc *TransactionService) ViewWalletTransaction(ctx context.Context, user entity.User) (result []entity.Transaction, err error) {
	// get wallet.
	wallet, err := uc.WalletRepo.GetWalletByUserId(ctx, user.CustomerXid)
	if helper.IsStructEmpty(wallet) || wallet.Status == helper.ConstantDisabled {
		return result, errors.New("wallet not found or wallet is disabled")
	}
	if err != nil {
		return result, err
	}

	result, err = uc.TransactionRepo.GetTransactionByWalletId(wallet.WalletId)

	return result, err
}
