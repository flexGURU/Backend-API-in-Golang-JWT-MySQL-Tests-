package user

import (
	"fmt"
	"net/http"

	"github.com/flexGURU/goAPI/auth"
	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
	"github.com/go-playground/validator/v10"
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

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)  {
        // Parse JSON payload
        var userloginPayload types.LoginUserPayload
        if err := utils.ParseJSON(r, &userloginPayload); err != nil {
            utils.WriteError(w, http.StatusBadRequest, err)
            return
        }
    
        // Validate the payload
        if err := utils.Validate.Struct(userloginPayload); err != nil {
            errors := err.(validator.ValidationErrors)
            utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload &%s", errors))
            return
        }

        user, err := h.store.GetUserByEmail(userloginPayload.Email)
        if err != nil {
            utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("user not found" ))

            return
        }

        if !auth.ComparePassword([]byte(user.Password), []byte(userloginPayload.Password)) {
            utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("email or password in incorrect"))
            return
        }

        utils.WriteJSON(w,http.StatusOK, map[string]string{ "token" : "" } )

}


func (h *Handler) handleregister(w http.ResponseWriter, r *http.Request) {
    // Parse JSON payload
    var userRegisterPayload types.RegisterUserPayload
    if err := utils.ParseJSON(r, &userRegisterPayload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // Validate the payload
    if err := utils.Validate.Struct(userRegisterPayload); err != nil {
        errors := err.(validator.ValidationErrors)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload &%s", errors))
        return
    }


    // Check if user exists
    user, err := h.store.GetUserByEmail(userRegisterPayload.Email)
    if err == nil  {
        utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("user with email: %v already exits", user.Email))
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
        utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("problem creating user %v ", err))
        return
    }

    utils.WriteJSON(w, http.StatusCreated, nil)
}

