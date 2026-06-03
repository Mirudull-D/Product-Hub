package cart

import (
	"Product-Hub/types"
	"Product-Hub/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.HandleCartCheckout).Methods(http.MethodPost)
}
func (h *Handler) HandleCartCheckout(w http.ResponseWriter, r *http.Request) {
	var payload types.CartCheckoutPayload

	ctx := r.Context()

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload &v", errors))
		return
	}

	productIds, err := GetCartitemsIds(payload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	ps, err := h.productStore.GetProductsById(ctx, productIds)
}
