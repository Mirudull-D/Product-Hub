package products

import (
	"Product-Hub/types"
	"Product-Hub/utils"
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
	router.HandleFunc("/product", h.HandleGetProduct).Methods(http.MethodGet)
}
func (h *Handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, products)
}
