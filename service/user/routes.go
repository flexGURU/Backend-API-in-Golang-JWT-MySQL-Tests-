package user

import (
	"log"
	"net/http"

	"github.com/flexGURU/goAPI/types"
	"github.com/flexGURU/goAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}

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
	log.Fatal("promlem unmarshalling json",err)
}
 err := utils.WriteJSON(w, userRegisterPayload) 
 if err != nil {
	log.Fatal(err)
 }


// check if User exits
	// if user doesn't create a new user

	
}