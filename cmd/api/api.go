package api

import (
	"Product-Hub/service/cart"
	"Product-Hub/service/order"
	"Product-Hub/service/products"
	"Product-Hub/service/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type application struct {
	addr string
	db   *sql.DB
	rdb  *redis.Client
}

func NewApplication(addr string, db *sql.DB, rdb *redis.Client) *application {
	return &application{
		addr: addr,
		db:   db,
		rdb:  rdb,
	}
}

func (app *application) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(app.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	productStore := products.NewStore(app.db, app.rdb)
	productService := products.NewHandler(productStore)
	productService.RegisterRoutes(subrouter)

	orderStore := order.NewStore(app.db)
	cartService := cart.NewHandler(orderStore, productStore, userStore)
	cartService.RegisterRoutes(subrouter)

	log.Print("Listening on ", app.addr)
	return http.ListenAndServe(app.addr, router)
}
