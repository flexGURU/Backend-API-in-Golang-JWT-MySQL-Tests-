package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/flexGURU/goAPI/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(add string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: add,
		db:   db,
	}

}

func (server *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	log.Println("listening on Address", server.addr)
	userHandler := user.NewHandler()
	userHandler.RegisterRoute(subrouter)

	return http.ListenAndServe(server.addr, router)

}