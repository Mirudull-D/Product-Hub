package main

import (
	"Product-Hub/cmd/api"
	"Product-Hub/config"
	db2 "Product-Hub/db"
	"database/sql"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {

	db, err := db2.NewPostgreSqlStorage(config.Envs.ConnString)

	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.0.108:6379",
	})

	app := api.NewApplication(config.Envs.Port, db, rdb)

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
