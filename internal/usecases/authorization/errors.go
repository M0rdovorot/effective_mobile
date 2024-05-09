package authorization

import (
	"fmt"
	"net/http"
)

type ErrorAuthorization struct {
	StatusCode int
	Msg string
}

func (e ErrorAuthorization) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrUnauthorized = ErrorAuthorization{
		StatusCode: http.StatusUnauthorized,
		Msg: `{"error": "unauthorized"}`,
	}
	ErrForbidden = ErrorAuthorization{
		StatusCode: http.StatusForbidden,
		Msg: `{"error": "forbidden"}`,
	}
)