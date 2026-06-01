package user

import (
	"Product-Hub/db/generated"
	"Product-Hub/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "wegh",
			LastName:  "ghjk,",
			Email:     "anana",
			Password:  "6789ijb",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected code %d got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(ctx context.Context, email string) (*generated.User, error) {

	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(ctx context.Context, id int) (*generated.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(ctx context.Context, payload types.RegisterUserPayload) error {
	return nil
}
