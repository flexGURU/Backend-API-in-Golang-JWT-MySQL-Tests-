package main

import (
	"log"

	"github.com/flexGURU/goAPI/cmd/api"
	"github.com/flexGURU/goAPI/config"
	"github.com/flexGURU/goAPI/db"
)

func main() {

	

	db, err := db.NewPostgre(&config.Envs)
	if err != nil {
		log.Fatal(err)
	}


	server := api.NewAPIServer(":8080",db  )

	if err := server.Run(); err != nil {
		log.Fatal("problem",err)
	}



}