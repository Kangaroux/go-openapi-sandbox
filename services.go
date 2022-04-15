package api

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UserService interface {
	Get(id int64) (*User, error)
}

type DBUserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &DBUserService{db: db}
}

func (s *DBUserService) Get(id int64) (*User, error) {
	u := User{}

	if err := s.db.Get(&u, "SELECT * FROM users WHERE id=$1", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}
