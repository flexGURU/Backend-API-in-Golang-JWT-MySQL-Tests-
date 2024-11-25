package user

import (
	"net/http"

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

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)  {

	
}

func (h *Handler) handleregister(w http.ResponseWriter, r *http.Request)  {

// get JSON Payload
var userRegisterPayload types.RegisterUserPayload

if err := utils.ParseJSON(r, &userRegisterPayload); err != nil {
	utils.WriteError(w, http.StatusBadGateway, err)
}

err := utils.WriteJSON(w, http.StatusAccepted, userRegisterPayload)
if err != nil {
	utils.WriteError(w, http.StatusBadRequest, err)
}
 


// check if User exits
	// if user doesn't create a new user

	
}