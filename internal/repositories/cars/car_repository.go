package cars

import (
	model "github.com/M0rdovorot/effective_mobile/internal/model"
)

type CarRepository interface{
	GetCars(filter map[string]any, page int) ([]model.Car, error)
	CreateCars(cars []model.Car) ([]int, error)
	PatchCar(id int, patchMap map[string]any) (error)
	DeleteCar(carId int) (error) 
}