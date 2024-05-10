package usecases

import (
	"context"
	"strconv"

	"github.com/M0rdovorot/effective_mobile/internal/model"
	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
)

func (usecases *Usecases) GetCars(ctx context.Context) ([]model.Car, error) {
	vars := ctxusecase.GetQueryVars(ctx)
	if vars == nil {
		return []model.Car{}, ErrNoVars
	}


	filter := map[string]any{}
	
	if regNum, ok := vars["regNum"]; ok{
		filter["regnum"] = regNum
	}
	if mark, ok := vars["mark"]; ok{
		filter["mark"] = mark
	}
	if model, ok := vars["model"]; ok{
		filter["model"] = model
	}
	if _, ok := vars["year"]; ok{
		year, err := strconv.Atoi(vars["year"])
		if err != nil {
			return []model.Car{}, ErrBadYear
		}
		filter["year"] = year
	}
	if name, ok := vars["name"]; ok{
		filter["name"] = name
	}
	if surname, ok := vars["surname"]; ok{
		filter["surname"] = surname
	}
	if patronymic, ok := vars["patronymic"]; ok{
		filter["patronymic"] = patronymic
	}

	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		page = 1
	}
	if page < 1 { 
		return []model.Car{}, ErrBadPage
	}
	
	cars, err := usecases.Cars.GetCars(filter, page)
	if err != nil {
		return []model.Car{}, err
	}

	for i := range cars {
		cars[i].Patronymic = cars[i].NullPatronymic.String
		cars[i].Year = int(cars[i].NullYear.Int16)
		cars[i].Owner.Name = cars[i].Name
		cars[i].Owner.Surname = cars[i].Surname
		cars[i].Owner.Patronymic = cars[i].Patronymic
	}

	return cars, nil
}