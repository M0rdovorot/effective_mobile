package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/db/postgresql"
	"github.com/M0rdovorot/effective_mobile/internal/handlers"
	"github.com/M0rdovorot/effective_mobile/internal/repositories/banners"
	"github.com/gorilla/mux"
)

func main() {
	pool := banners.NewPool(fmt.Sprintf("%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	
	var db postgresql.Database
	err := db.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	
	cashRepository := banners.CreateBannerRedisStorage(pool)
	bannersRepository := banners.CreateBannerStorage(db.GetDB())
	bannerHandler := handlers.CreateBannerHandler(bannersRepository, cashRepository)
	
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user_banner", handlers.NewWrapper(bannerHandler.GetUserBanner).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/banner", handlers.NewWrapper(bannerHandler.GetAllBanners).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/banner", handlers.NewWrapper(bannerHandler.CreateBanner).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/banner/{id:[0-9]+}", handlers.NewWrapper(bannerHandler.PatchBanner).ServeHTTP).Methods(http.MethodPatch)
	router.HandleFunc("/api/v1/banner/{id:[0-9]+}", handlers.NewWrapper(bannerHandler.DeleteBanner).ServeHTTP).Methods(http.MethodDelete)

	err = http.ListenAndServe(":8001", router)
	log.Println(err)
}