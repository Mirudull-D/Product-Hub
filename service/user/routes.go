package user

import (
	"Product-Hub/service/auth"
	"Product-Hub/types"
	"Product-Hub/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

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
	_, err := h.store.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest,
			fmt.Errorf("user already exists with email %s", payload.Email))
		return
	}
	hashedPassword, err := auth.Hashpassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	payload.Password = hashedPassword

	err = h.store.CreateUser(ctx, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, nil)

}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

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
	user, err := h.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found,invalid email or password"))
		return
	}
	if !auth.Comparepasswords(user.Password, payload.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found,invalid email or password"))
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{"token": ""})
}
