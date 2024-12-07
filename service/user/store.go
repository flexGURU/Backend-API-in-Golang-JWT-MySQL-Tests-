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
    var user types.User
    err := s.db.QueryRow("SELECT firstname, lastname, email  FROM users WHERE email = $1", email).
        Scan(&user.FirstName, &user.LastName, &user.Email)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("error fetching user: %v", err)
    }

    return &user, nil
}



func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

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