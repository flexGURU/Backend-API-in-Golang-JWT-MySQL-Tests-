 package user

import (
	"database/sql"
	"fmt"

	"github.com/flexGURU/goAPI/types"
)

type Store struct {
	db *sql.DB
}


func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email = ? ", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = ScanRows(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
	
}

func (s *Store) CreateUser(types.User) ( error) {

	return nil	
}

func ScanRows(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}