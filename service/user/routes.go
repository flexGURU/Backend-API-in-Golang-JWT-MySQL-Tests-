package user

import (
	"fmt"
	"net/http"

	"github.com/flexGURU/goAPI/auth"
	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
	"github.com/gorilla/mux"
)


type Handler struct { 
	store types.UserStore
	
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}

}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleregister).Methods("POST")


}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)  {}

func (h *Handler) handleregister(w http.ResponseWriter, r *http.Request) {
    // Parse JSON payload
    var userRegisterPayload types.RegisterUserPayload
    if err := utils.ParseJSON(r, &userRegisterPayload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("jameni %s", err))
        return
    }

    // Check if user exists
    user, err := h.store.GetUserByEmail(userRegisterPayload.Email)
    if err == nil && user != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this email %s exists", user.Email))
        return
    }

    // Hash password
    hashedPwd, err := auth.HashPassword(userRegisterPayload.Password)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // Create new user
    err = h.store.CreateUser(types.User{
        FirstName: userRegisterPayload.FirstName,
        LastName:  userRegisterPayload.LastName,
        Email:     userRegisterPayload.Email,
        Password:  hashedPwd,
    })
    if err != nil {
        utils.WriteError(w, http.StatusBadGateway, err)
        return
    }

    utils.WriteJSON(w, http.StatusCreated, nil)
}
