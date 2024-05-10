package usecases

import (
	"context"
	"strconv"

	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
)

func (usecases *Usecases) DeleteCar(ctx context.Context) (error) {
	vars := ctxusecase.GetVars(ctx)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return ErrBadID
	}

	return usecases.Cars.DeleteCar(id)
}