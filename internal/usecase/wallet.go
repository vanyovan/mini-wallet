package usecase

import (
	"context"
	"errors"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
)

type WalletService struct {
	UserRepo   repo.UserRepo
	WalletRepo repo.WalletRepo
}

type WalletServiceProvider interface {
	CreateEnabledWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error)
}

func NewWalletService(UserRepo repo.UserRepo, WalletRepo repo.WalletRepo) WalletService {
	return WalletService{
		UserRepo:   UserRepo,
		WalletRepo: WalletRepo,
	}
}

func (uc *WalletService) CreateEnabledWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error) {
	// get wallet. if there's wallet by this user, cannot create new wallet
	result, err = uc.WalletRepo.GetWalletByUserId(ctx, param.CustomerXid)
	if !helper.IsStructEmpty(result) {
		return result, errors.New("wallet already exists. you dont need to enabled it again")
	}

	//create new wallet
	result, err = uc.WalletRepo.CreateWallet(ctx, param.CustomerXid)
	return result, err
}
