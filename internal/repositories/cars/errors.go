package cars

import (
	"fmt"
	"net/http"
)

type ErrorCars struct {
	StatusCode int
	Msg string
}

func (e ErrorCars) Error() string {
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrNotFound = ErrorCars{
		StatusCode: http.StatusNotFound,
		Msg: `{"error": "not found"}`,
	}

	ErrResultNotOk = ErrorCars{
		StatusCode: http.StatusInternalServerError,
		Msg: `{"error": "not ok"}`,
	} 
)