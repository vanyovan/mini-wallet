package errors

import (
	"fmt"
	"net/http"
)

type HandleError struct {
	HttpStatus int
	Data       interface{}
}

func (e HandleError) Error() string {
	return fmt.Sprintf("%s", e.Data)
}

func ValidationError(data interface{}) HandleError {
	return HandleError{
		HttpStatus: http.StatusBadRequest,
		Data:       data,
	}
}

func NotFound(data interface{}) HandleError {
	return HandleError{
		HttpStatus: http.StatusNotFound,
		Data:       data,
	}
}

func UnauthorizedError(data interface{}) HandleError {
	return HandleError{
		HttpStatus: http.StatusUnauthorized,
		Data:       data,
	}
}
