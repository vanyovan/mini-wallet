package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

type Handler struct {
	UserUsecase usecase.UserService
}

type TokenUsecase interface {
	HandleInitWallet(w http.ResponseWriter, r *http.Request)
}

func NewHandler(UserUsecase usecase.UserService) *Handler {
	return &Handler{
		UserUsecase: UserUsecase,
	}
}

func (h *Handler) HandleInitWallet(w http.ResponseWriter, r *http.Request) {
	request := InitWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	if request.CustomerXid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	param := entity.WalletRequestParam{
		CustomerXid: request.CustomerXid,
	}
	result, err := h.UserUsecase.CreateUser(context.TODO(), param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	jsonResponse, err := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
