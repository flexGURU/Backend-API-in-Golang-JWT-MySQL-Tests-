package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/flexGURU/goAPI/config"
	_ "github.com/lib/pq"
)


func NewPostgre(cfg *config.Config) (*sql.DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
			cfg.DBUser, 
			cfg.DBPassword, 
			cfg.DBHost, 
			cfg.Port, 
			cfg.DBName,
		)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	 
	return db, nil

}