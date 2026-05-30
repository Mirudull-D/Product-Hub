package main

import (
	"Product-Hub/cmd/api"
	"log"
)

func main() {

	app := api.NewApplication(":8080", nil)
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
