package handlers

import (
	"fmt"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/internal/repositories/banners"
	"github.com/M0rdovorot/effective_mobile/internal/usecases/authorization"
)

type ErrorHttp struct{
	StatusCode int
	Msg string
}

func (e ErrorHttp) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrNoVars = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg: `{"error": "no vars"}`,
	}
	
	ErrBadTagID = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad tag id"}`,
	}

	ErrBadID = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad id"}`,
	}

	ErrBadFeatureID = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad feature id"}`,
	}

	ErrDataBase = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"db"}`,
	}

	ErrEncoding = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"encoding_json"}`,
	}

	ErrDecoding = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"wrong_json"}`,
	}

	ErrUnauthorized = ErrorHttp{
		StatusCode: http.StatusUnauthorized,
		Msg:        `{"error":"unauthorized"}`,
	}
)

func WriteHttpError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case ErrorHttp:
		err := err.(ErrorHttp)
		http.Error(w, err.Msg, err.StatusCode)
	case authorization.ErrorAuthorization:
		err := err.(authorization.ErrorAuthorization)
		http.Error(w, err.Msg, err.StatusCode)
	case banners.ErrorBanners:
		err := err.(banners.ErrorBanners)
		http.Error(w, err.Msg, err.StatusCode)
	default:
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
	}
}