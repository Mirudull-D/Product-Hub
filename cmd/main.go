package main

import (
	"Product-Hub/cmd/api"
	"Product-Hub/config"
	db2 "Product-Hub/db"
	"log"
)

func main() {

	db, err := db2.NewPostgreSqlStorage(config.Envs.ConnString)

	app := api.NewApplication(config.Envs.Port, db)
	err2 := app.Start()
	if err2 != nil {
		log.Fatal(err)
	}
}
