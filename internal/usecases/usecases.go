package usecases

import (
	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/internal/repositories/cars"
)

type Usecases struct {
	Cars cars.CarRepository
	config *configs.Config
}

func CreateUsecases(cars cars.CarRepository, config *configs.Config) *Usecases {
	return &Usecases{
		Cars: cars,
		config: config,
	}
}