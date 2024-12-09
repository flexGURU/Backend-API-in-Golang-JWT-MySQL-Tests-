package product

import (
	"database/sql"

	"github.com/flexGURU/goAPI/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store  {
	return &Store{db: db}
	
}

func (s *Store) GetProducts() ([]types.Product, error)  {

	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := ScanRows(rows)
		if err != nil {
			return nil, err
		}
		
		products = append(products, *p)

	}

	return products, nil
	
}


func ScanRows(rows *sql.Rows) (*types.Product, error) {
	products := new(types.Product)

	err := rows.Scan(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}