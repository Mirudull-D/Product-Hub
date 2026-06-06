package user

import (
	"Product-Hub/config"
	"Product-Hub/service/auth"
	"Product-Hub/types"
	"Product-Hub/utils"
	"fmt"
	"net/http"
	"time"

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
	router.HandleFunc("/health", h.HandleHealth).Methods(http.MethodGet)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		fmt.Println("Login total:", time.Since(start))
	}()
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
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJwt(secret, int(user.ID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !auth.Comparepasswords(user.Password, payload.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found,invalid email or password"))
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{"token": token})
}

func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	err := h.store.PingDb(r.Context())
	if err != nil {
		err = utils.WriteJson(w, http.StatusServiceUnavailable,
			map[string]string{
				"status": "database down",
			},
		)
		if err != nil {
			return
		}
		return
	}

	err = utils.WriteJson(w, http.StatusOK,
		map[string]string{
			"status": "ok",
		},
	)
	if err != nil {
		return
	}
}
