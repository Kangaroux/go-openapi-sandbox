package api

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserService interface {
	Get(id int64) (*User, error)
	Exists(where string, args ...interface{}) (bool, error)
}

type DBUserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &DBUserService{db: db}
}

func (s *DBUserService) Exists(where string, args ...interface{}) (bool, error) {
	count := 0
	q := "SELECT COUNT(*) FROM users WHERE " + where

	if err := s.db.Get(&count, q, args...); err != nil {
		log.Print(err)
		return false, err
	}

	return count > 0, nil
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
