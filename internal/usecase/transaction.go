package usecase

import (
	"github.com/vanyovan/mini-wallet.git/internal/repo"
)

type TransactionService struct {
	UserRepo        repo.UserRepo
	WalletRepo      repo.WalletRepo
	TransactionRepo repo.TransactionRepo
}

type TransactionServiceProvider interface {
}

func NewTransactionService(UserRepo repo.UserRepo, WalletRepo repo.WalletRepo, TransactionRepo repo.TransactionRepo) TransactionService {
	return TransactionService{
		UserRepo:        UserRepo,
		WalletRepo:      WalletRepo,
		TransactionRepo: TransactionRepo,
	}
}
