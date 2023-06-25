package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/usecase/mock_usecase"
)

func TestUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	any := gomock.Any()

	userRepo := mock_usecase.NewMockUserServiceProvider(ctrl)

	service := NewUserService(userRepo)

	tests := []struct {
		name    string
		param   entity.UserRequestParam
		want    entity.UserResponse
		wantErr bool
		mock    func()
	}{
		{"failed_create_user_exist", entity.UserRequestParam{CustomerXid: ""},
			entity.UserResponse{},
			true,
			func() {
				userRepo.EXPECT().GetUserByUserId(any).Return(entity.User{CustomerXid: "123"}, nil)
			}},
		{"failed_create_user_error_getting_user", entity.UserRequestParam{CustomerXid: ""},
			entity.UserResponse{},
			true,
			func() {
				userRepo.EXPECT().GetUserByUserId(any).Return(entity.User{}, errors.New("error"))
			}},
		{"success_create_user", entity.UserRequestParam{CustomerXid: "123"},
			entity.UserResponse{Token: "123"},
			false,
			func() {
				userRepo.EXPECT().GetUserByUserId(any).Return(entity.User{}, nil)
				userRepo.EXPECT().CreateUser(any, any).Return("123", nil)
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			_, err := service.CreateUser(context.TODO(), tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.SetOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
