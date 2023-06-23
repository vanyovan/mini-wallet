package usecase

import (
	"context"
	"errors"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/repo"
)

type UserService struct {
	UserRepo repo.UserRepo
}

type UserServiceProvider interface {
	CreateUser(ctx context.Context, param entity.UserRequestParam) (result entity.UserResponse, err error)
	GetUserByToken(ctx context.Context, token string) (entity.User, error)
}

func NewUserService(UserRepo repo.UserRepo) UserService {
	return UserService{
		UserRepo: UserRepo,
	}
}

func (uc *UserService) CreateUser(ctx context.Context, param entity.UserRequestParam) (result entity.UserResponse, err error) {
	// find user, if user already exists return error
	user, err := uc.UserRepo.GetUserByUserId(param.CustomerXid)
	if !helper.IsStructEmpty(user) {
		return result, errors.New("user already exists")
	}

	// add database user and token
	token, err := uc.UserRepo.CreateUser(ctx, param.CustomerXid)
	result.Token = token
	return result, err
}

func (uc *UserService) GetUserByToken(ctx context.Context, token string) (entity.User, error) {
	currentUser, err := uc.UserRepo.GetUserByToken(ctx, token)
	if err != nil {
		return entity.User{}, err
	}

	return currentUser, nil

}
