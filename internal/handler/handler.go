package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
	"github.com/vanyovan/mini-wallet.git/internal/helper"
	"github.com/vanyovan/mini-wallet.git/internal/usecase"
)

type Handler struct {
	UserUsecase        usecase.UserService
	WalletUsecase      usecase.WalletService
	TransactionUsecase usecase.TransactionService
}

type TokenUsecase interface {
	HandleInitWallet(w http.ResponseWriter, r *http.Request)
}

func NewHandler(UserUsecase usecase.UserService, WalletUsecase usecase.WalletService, TransactionUsecase usecase.TransactionService) *Handler {
	return &Handler{
		UserUsecase:        UserUsecase,
		WalletUsecase:      WalletUsecase,
		TransactionUsecase: TransactionUsecase,
	}
}

func (h *Handler) HandleInitWallet(w http.ResponseWriter, r *http.Request) {
	request := InitWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	if request.CustomerXid == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return

	}

	param := entity.UserRequestParam{
		CustomerXid: request.CustomerXid,
	}

	result, err := h.UserUsecase.CreateUser(context.TODO(), param)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleEnableWallet(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	result, err := h.WalletUsecase.CreateEnableWallet(r.Context(), currentUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleViewWallet(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	result, err := h.WalletUsecase.ViewWallet(r.Context(), currentUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleViewTransaction(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return

	}

	result, err := h.TransactionUsecase.ViewWalletTransaction(r.Context(), currentUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"transactions": result,
		},
	})
}

func (h *Handler) HandleDepositWallet(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return

	}

	request := TransactionRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	param := entity.TransactionRequest{
		Amount:      request.Amount,
		ReferenceId: request.ReferenceId,
	}

	result, err := h.TransactionUsecase.CreateDepositWallet(r.Context(), currentUser, param)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleWithdrawalWallet(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	request := TransactionRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	param := entity.TransactionRequest{
		Amount:      request.Amount,
		ReferenceId: request.ReferenceId,
	}

	result, err := h.TransactionUsecase.CreateWithdrawalWallet(r.Context(), currentUser, param)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleDisableWallet(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helper.FromContext(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	result, err := h.WalletUsecase.CreateDisableWallet(r.Context(), currentUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}
