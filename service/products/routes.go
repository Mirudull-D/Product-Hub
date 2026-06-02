package products

import (
	"Product-Hub/types"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", HandleGetProduct).Methods(http.MethodGet)
}
func HandleGetProduct(w http.ResponseWriter, r *http.Request) {}
