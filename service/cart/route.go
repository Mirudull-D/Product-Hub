package cart

import (
	"Product-Hub/service/auth"
	"Product-Hub/types"
	"Product-Hub/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
	db           *sql.DB
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJwtAuth(h.HandleCartCheckout, h.userStore)).Methods(http.MethodPost)
}
func (h *Handler) HandleCartCheckout(w http.ResponseWriter, r *http.Request) {
	var payload types.CartCheckoutPayload
	ctx := r.Context()
	userId, err2 := auth.GetUserIdfromRequest(ctx)
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, err2)
		return
	}

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
		return
	}
	ps, err := h.productStore.GetProductsById(ctx, productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	orderId, totalPrice, err := h.createOrder(ctx, ps, payload.Items, userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJson(w, http.StatusCreated, map[string]any{
		"orderId":    orderId,
		"totalPrice": totalPrice,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
