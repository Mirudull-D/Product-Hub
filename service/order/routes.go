package order

import (
	"Product-Hub/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{
		store: store,
	}

}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	//router.HandleFunc("/order",).Methods(http.MethodGet)
}
