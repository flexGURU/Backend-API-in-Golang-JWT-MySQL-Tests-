package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flexGURU/goAPI/types"
	"github.com/gorilla/mux"
)


type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	// Mock implementation returning nil user and nil error
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(types.User) (error)  {
	return nil
	
}

func TestServiceHandlers(t *testing.T) {
    userStore := &mockUserStore{}
    handler := NewHandler(userStore)

    t.Run("should fail if user payload is invalid", func(t *testing.T) {
        payload := types.RegisterUserPayload{
            FirstName: "mukuna",
            LastName:  "john",
            Email:     "mukunajohn329@gmail.com", // Invalid email
            Password:  "123",
        }
        marshalledPayload, _ := json.Marshal(payload)

        req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
        router := mux.NewRouter()
        router.HandleFunc("/register", handler.handleregister)

        router.ServeHTTP(rr, req)

        if rr.Code != http.StatusBadRequest {
            t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rr.Code)
        }
    })
}
