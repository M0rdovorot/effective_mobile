package banners

import (
	"fmt"
	"net/http"
)

type ErrorBanners struct {
	StatusCode int
	Msg string
}

func (e ErrorBanners) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrNotFound = ErrorBanners{
		StatusCode: http.StatusNotFound,
		Msg: `{"error": "not found"}`,
	}

	ErrForbidden = ErrorBanners{
		StatusCode: http.StatusForbidden,
		Msg: `{"error": "forbidden"}`,
	}

	ErrResultNotOk = ErrorBanners{
		StatusCode: http.StatusInternalServerError,
		Msg: `{"error": "not ok"}`,
	} 
)