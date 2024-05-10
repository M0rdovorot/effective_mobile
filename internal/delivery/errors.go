package handlers

import (
	"fmt"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/internal/repositories/cars"
)

type ErrorHttp struct{
	StatusCode int
	Msg string
}

func (e ErrorHttp) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrEncoding = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"encoding_json"}`,
	}

	ErrDecoding = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"wrong_json"}`,
	}
)

func WriteHttpError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case ErrorHttp:
		err := err.(ErrorHttp)
		http.Error(w, err.Msg, err.StatusCode)
	case cars.ErrorCars:
		err := err.(cars.ErrorCars)
		http.Error(w, err.Msg, err.StatusCode)
	default:
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
	}
}