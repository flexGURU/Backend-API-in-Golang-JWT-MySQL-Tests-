package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/flexGURU/goAPI/types"
)

type Store struct {
	db *sql.DB
}

// UpdateProduct implements types.ProductStore.
func (s *Store) UpdateProduct(types.Product) error {
	panic("unimplemented")
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}

}

func (s *Store) GetProducts() ([]types.Product, error) {

	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error reteriveing data", err)
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

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := make([]string, len(productIDs))

	args := make([]interface{}, len(productIDs))

	for i, value := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = value
	}

	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%S)", strings.Join(placeholders, ","))

	rows, err := s.db.Query(query, args)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
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

	err := rows.Scan(
		&products.ID,
		&products.Name,
		&products.Description,
		&products.Image,
		&products.Price,
		&products.Quantity,
		&products.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error reteriveing data", err)
	}

	return products, nil
}
