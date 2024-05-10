package handlers

import (
	"context"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/internal/model"
	"github.com/M0rdovorot/effective_mobile/internal/usecases"
)

type CarHandler struct {
	Usecases *usecases.Usecases
}

func CreateCarHandler(usecases *usecases.Usecases) *CarHandler {
	return &CarHandler{
		Usecases: usecases,
	}
}

func (handler *CarHandler) GetCars(ctx context.Context, form EmptyForm) ([]model.Car, int, error) {
	cars, err := handler.Usecases.GetCars(ctx)
	if err != nil {
		return []model.Car{}, 0, err
	}

	for i := range cars {
		cars[i].Owner.Name = cars[i].Name
		cars[i].Owner.Surname = cars[i].Surname
		cars[i].Owner.Patronymic = cars[i].Patronymic
	}

	return cars, http.StatusOK, nil
}

func (handler *CarHandler) CreateCars(ctx context.Context, regNums RegNumsForm) (IDs, int, error) {
	ids, err := handler.Usecases.CreateCars(ctx, regNums.RegNums)
	if err != nil {
		return IDs{}, 0, err
	}

	return IDs{IDs: ids}, http.StatusCreated, nil
}

func (handler *CarHandler) PatchCar(ctx context.Context, patchMap PatchMapForm) (any, int, error) {
	return 0, http.StatusNoContent, handler.Usecases.PatchCar(ctx, patchMap)
}

func (handler *CarHandler) DeleteCar(ctx context.Context, form EmptyForm) (any, int, error) {
	return 0, http.StatusNoContent, handler.Usecases.DeleteCar(ctx)
}
