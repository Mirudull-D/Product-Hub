package api

import (
	"Product-Hub/service/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type application struct {
	addr string
	db   *sql.DB
}

func NewApplication(addr string, db *sql.DB) *application {
	return &application{addr: addr, db: db}
}

func (app *application) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter)

	log.Print("Listening on ", app.addr)
	return http.ListenAndServe(app.addr, nil)
}
