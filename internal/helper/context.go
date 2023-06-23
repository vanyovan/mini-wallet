package helper

import (
	"context"
	"errors"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
)

func Inject(ctx context.Context, data entity.User) context.Context {
	return context.WithValue(ctx, userKey, data)
}

func FromContext(ctx context.Context) (entity.User, error) {
	currentUser, ok := ctx.Value(userKey).(entity.User)
	if !ok {
		return entity.User{}, errors.New("client invalid")
	}
	if IsStructEmpty(currentUser) {
		return entity.User{}, errors.New("client invalid")
	}
	if currentUser.CustomerXid == "" {
		return entity.User{}, errors.New("client invalid")
	}
	return currentUser, nil
}
