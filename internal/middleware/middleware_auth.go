package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

func AuthenticateUser(userService usecase.UserServiceProvider) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := strings.Split(r.Header.Get("Authorization"), "Token ")
			if len(authorization) == 0 {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "fail",
					"data": map[string]interface{}{
						"error": errors.New("token invalid").Error(),
					},
				})
				return
			}

			token := authorization[1]

			currentUser, err := userService.GetUserByToken(r.Context(), token)
			if helper.IsStructEmpty(currentUser) || err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "fail",
					"data": map[string]interface{}{
						"error": errors.New("token invalid").Error(),
					},
				})
				return
			}

			ctx := helper.Inject(r.Context(), currentUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
