package usecases

import (
	"context"
	"strconv"

	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
)

func (usecases *Usecases) PatchCar(ctx context.Context, patchMap map[string]any) error {
	vars := ctxusecase.GetVars(ctx)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return ErrBadID
	}

	filter, err := validateCarMap(patchMap)
	if err != nil {
		return err
	}

	return usecases.Cars.PatchCar(id, filter) 
}