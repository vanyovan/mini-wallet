package response

import (
	"encoding/json"
	"net/http"

	"github.com/vanyovan/mini-wallet.git/internal/helper/errors"
)

func Success(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func Failed(w http.ResponseWriter, err error) {
	isHandleError, handleError := handleError(err)
	if isHandleError {
		failedHandled(w, handleError.HttpStatus, handleError.Data)
		return
	}

	failedError(w, err)
}

func failedError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "error",
		"message": err.Error(),
	})
}

func failedHandled(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "fail",
		"data": map[string]interface{}{
			"error": data,
		},
	})
}

func handleError(err error) (bool, errors.HandleError) {
	isHandledError := false
	var definedError errors.HandleError

	switch err := err.(type) {
	case errors.HandleError:
		definedError = err
		isHandledError = true
	}

	return isHandledError, definedError
}
