package helper

import (
	"context"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
)

type key string

var clientKey = key("client_context")

// Inject the client into the current context
func Inject(ctx context.Context, data entity.User) context.Context {
	return context.WithValue(ctx, clientKey, data)
}
