package main

import (
	"log"

	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/db/postgresql"
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
	err = db.MigrateUp()
	if err != nil {
		log.Println(err)
	}
}
