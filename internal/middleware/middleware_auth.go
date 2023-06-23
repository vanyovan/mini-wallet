package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/helper/response"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

func AuthenticateUser(userService usecase.UserServiceProvider) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := strings.Split(r.Header.Get("Authorization"), "Token ")
			if len(authorization) == 0 {
				response.Failed(w, errors.New("token invalid"))
				return
			}

			token := authorization[1]

			currentUser, err := userService.GetUserByToken(r.Context(), token)
			if err != nil {
				response.Failed(w, err)
				return
			}
			if helper.IsStructEmpty(currentUser) {
				response.Failed(w, errors.New("token invalid"))
				return
			}

			ctx := helper.Inject(r.Context(), currentUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
