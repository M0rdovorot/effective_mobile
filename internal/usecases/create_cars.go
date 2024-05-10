package usecases

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/internal/model"
)

func makeCarInfoRequest(url string) (model.Car, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return model.Car{}, err
	}
	defer resp.Body.Close()

	var car model.Car
	err = json.NewDecoder(resp.Body).Decode(&car)
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (usecases *Usecases) CreateCars(ctx context.Context, regNums []string) ([]int, error) {
	if len(regNums) == 0 {
		return []int{}, ErrNoVars
	}

	var cars []model.Car
	for _, regNum := range regNums {
		car, err := makeCarInfoRequest(usecases.config.CarInfoAPI + "?regNum=" + regNum)
		if err != nil {
			log.Println(err)
			continue
		}
		cars = append(cars, car)
	}

	if len(cars) == 0 {
		return []int{}, ErrBadRegNums
	}
	ids, err := usecases.Cars.CreateCars(cars)
	if err != nil {
		return []int{}, err
	}

	return ids, nil
}
