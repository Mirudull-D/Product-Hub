package main

import (
	"Product-Hub/cmd/api"
	"Product-Hub/config"
	db2 "Product-Hub/db"
	"database/sql"
	"log"
)

func main() {

	db, err := db2.NewPostgreSqlStorage(config.Envs.ConnString)

	app := api.NewApplication(config.Envs.Port, db)

	initStorage(db)

	err2 := app.Start()
	if err2 != nil {
		log.Fatal(err)
	}
}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB connected Successfully ...!!")
}
