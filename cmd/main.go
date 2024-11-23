package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/flexGURU/goAPI/cmd/api"
	"github.com/flexGURU/goAPI/config"
	"github.com/flexGURU/goAPI/db"
)

func main() {

	db, _ := initDB()
		
	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*sql.DB, error) {
	db, err := db.NewPostgre(&config.Envs)
	if err != nil {
		log.Fatal(err)
	}
	 if err = db.Ping(); err != nil {
		log.Fatal(err)
	 }
	 fmt.Println("database connection succesful")
	 return db, nil
	
}

