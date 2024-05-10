package usecases

import (
	"fmt"
	"net/http"
)

type ErrorUsecases struct{
	StatusCode int
	Msg string
}

func (e ErrorUsecases) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrNoVars = ErrorUsecases{
		StatusCode: http.StatusBadRequest,
		Msg: `{"error": "no vars"}`,
	}

	ErrBadRegNums = ErrorUsecases{
		StatusCode: http.StatusBadRequest,
		Msg: `{"error": "bad regNums"}`,
	}

	ErrBadID = ErrorUsecases{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad id"}`,
	}

	ErrBadYear = ErrorUsecases{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad year"}`,
	}

	ErrBadPage = ErrorUsecases{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"bad page"}`,
	}
)