package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
	"github.com/vanyovan/mini-wallet.git/internal/repo/wrapper"
)

type TransactionService struct {
	UserRepo        repo.UserRepo
	WalletRepo      repo.WalletRepo
	TransactionRepo repo.TransactionRepo
	SqlWrapper      wrapper.SqlWrapper
}

type TransactionServiceProvider interface {
	CreateDepositWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error)
	ViewWalletTransaction(ctx context.Context, user entity.User) (result entity.Transaction, err error)
	CreateWithdrawalWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error)
}

func NewTransactionService(UserRepo repo.UserRepo, WalletRepo repo.WalletRepo, TransactionRepo repo.TransactionRepo, SqlWrapper wrapper.SqlWrapper) TransactionService {
	return TransactionService{
		UserRepo:        UserRepo,
		WalletRepo:      WalletRepo,
		TransactionRepo: TransactionRepo,
		SqlWrapper:      SqlWrapper,
	}
}

func (uc *TransactionService) CreateDepositWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error) {
	// get wallet
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

	//insert transaction with success status
	transactionId, err := uc.TransactionRepo.CreateWalletTransaction(ctx, wallet.WalletId, status, helper.ConstantDeposit, param.Amount, param.ReferenceId)
	if err != nil {
		status = helper.ConstantFailed
		return wallet, err
	}

	go func() {
		time.Sleep(5 * time.Second)
		uc.SqlWrapper.BeginTx(ctx, func(ctx context.Context) error {
			//update wallet amount
			err := uc.TransactionRepo.UpdateWalletBalanceByUserId(ctx, wallet.Balance, user.CustomerXid)
			if err != nil {
				status = helper.ConstantFailed

				//update transaction with failed status
				err := uc.TransactionRepo.UpdateTransactionStatusByTransactionId(ctx, status, transactionId)
				return err
			}
			return nil
		})
	}()

	return wallet, err
}

func (uc *TransactionService) CreateWithdrawalWallet(ctx context.Context, user entity.User, param entity.TransactionRequest) (result entity.Wallet, err error) {
	// get wallet
	wallet, err := uc.WalletRepo.GetWalletByUserId(ctx, user.CustomerXid)
	if helper.IsStructEmpty(wallet) || wallet.Status == helper.ConstantDisabled {
		return result, errors.New("wallet not found or wallet is disabled")
	}
	if err != nil {
		return result, err
	}

	var status = helper.ConstantSuccess

	if wallet.Balance < param.Amount {
		return result, errors.New("balance not sufficient")
	}

	//count balance
	wallet.Balance = wallet.Balance - param.Amount

	//insert transaction with success status
	transactionId, err := uc.TransactionRepo.CreateWalletTransaction(ctx, wallet.WalletId, status, helper.ConstantDeposit, param.Amount, param.ReferenceId)
	if err != nil {
		status = helper.ConstantFailed
		return wallet, err
	}

	go func() {
		time.Sleep(5 * time.Second)
		uc.SqlWrapper.BeginTx(ctx, func(ctx context.Context) error {
			//update wallet amount
			err := uc.TransactionRepo.UpdateWalletBalanceByUserId(ctx, wallet.Balance, user.CustomerXid)
			if err != nil {
				status = helper.ConstantFailed

				//update transaction with failed status
				err := uc.TransactionRepo.UpdateTransactionStatusByTransactionId(ctx, status, transactionId)
				return err
			}
			return nil
		})
	}()
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
