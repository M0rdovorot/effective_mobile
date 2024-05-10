package main

import (
	"log"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/db/postgresql"
	_ "github.com/M0rdovorot/effective_mobile/docs"
	handlers "github.com/M0rdovorot/effective_mobile/internal/delivery"
	"github.com/M0rdovorot/effective_mobile/internal/repositories/cars"
	"github.com/M0rdovorot/effective_mobile/internal/usecases"

	_ "github.com/go-swagger/go-swagger"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	conf := configs.New()

	var db postgresql.Database
	err := db.Connect(conf.DB)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	carsRepository := cars.CreateCarStorage(db.GetDB(), conf)
	carsUsecases := usecases.CreateUsecases(carsRepository, conf)
	carHandler := handlers.CreateCarHandler(carsUsecases)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/car", handlers.NewWrapper(carHandler.GetCars).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/car", handlers.NewWrapper(carHandler.CreateCars).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/car/{id:[0-9]+}", handlers.NewWrapper(carHandler.PatchCar).ServeHTTP).Methods(http.MethodPatch)
	router.HandleFunc("/api/v1/car/{id:[0-9]+}", handlers.NewWrapper(carHandler.DeleteCar).ServeHTTP).Methods(http.MethodDelete)

	err = http.ListenAndServe(conf.BackendServerPort, router)
	log.Println(err)
}
