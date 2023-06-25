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
	CreateEnableWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error)
	CreateDisableWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error)
	ViewWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error)
}

func NewWalletService(UserRepo repo.UserRepo, WalletRepo repo.WalletRepo) WalletService {
	return WalletService{
		UserRepo:   UserRepo,
		WalletRepo: WalletRepo,
	}
}

func (uc *WalletService) CreateEnableWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error) {
	// get wallet. if there's wallet by this user, cannot create new wallet
	result, err = uc.WalletRepo.GetWalletByUserId(ctx, param.CustomerXid)
	if !helper.IsStructEmpty(result) && result.Status == helper.ConstantEnabled {
		return result, errors.New("already enabled")
	}

	if result.Status == helper.ConstantDisabled {
		//has to update the wallet to enable
		updatedAt, err := uc.WalletRepo.UpdateStatusByUserId(ctx, helper.ConstantEnabled, param.CustomerXid)
		if err != nil {
			return result, err
		}
		result.Status = helper.ConstantEnabled
		result.EnabledAt = &updatedAt
	} else {
		//create new wallet
		result, err = uc.WalletRepo.CreateWallet(ctx, param.CustomerXid)
		if err != nil {
			return result, err
		}
	}
	return result, err
}

func (uc *WalletService) CreateDisableWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error) {
	// get wallet.
	result, err = uc.WalletRepo.GetWalletByUserId(ctx, param.CustomerXid)
	if helper.IsStructEmpty(result) || result.Status == helper.ConstantDisabled {
		return result, errors.New("wallet not found or wallet already disabled")
	}

	if err != nil {
		return result, err
	}

	//disable wallet
	updatedAt, err := uc.WalletRepo.UpdateStatusByUserId(ctx, helper.ConstantDisabled, param.CustomerXid)
	if err != nil {
		return result, err
	}

	result.Status = helper.ConstantDisabled
	result.DisabledAt = &updatedAt

	return result, err
}

func (uc *WalletService) ViewWallet(ctx context.Context, param entity.User) (result entity.Wallet, err error) {
	result, err = uc.WalletRepo.GetWalletByUserId(ctx, param.CustomerXid)
	if helper.IsStructEmpty(result) || result.Status == helper.ConstantDisabled {
		return result, errors.New("wallet disabled")
	}

	return result, err
}
