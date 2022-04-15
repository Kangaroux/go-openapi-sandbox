package api

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserService interface {
	Create(u *User) error
	Count(where string, args ...interface{}) (int, error)
	Exists(where string, args ...interface{}) (bool, error)
	Get(where string, args ...interface{}) (*User, error)
	List(query string, args ...interface{}) ([]User, error)
}

type DBUserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &DBUserService{db: db}
}

func (s *DBUserService) Create(u *User) error {
	q := "INSERT INTO users (updated_at, username, email, password) " +
		"VALUES (:updated_at, :username, :email, :password) RETURNING *"

	u.UpdatedAt = time.Now()
	result, err := s.db.NamedQuery(q, u)

	if err != nil {
		return err
	}

	defer result.Close()

	result.Next()
	result.StructScan(&u)

	return nil
}

func (s *DBUserService) Count(where string, args ...interface{}) (int, error) {
	count := 0
	q := "SELECT COUNT(*) FROM users WHERE " + where

	if err := s.db.Get(&count, q, args...); err != nil {
		log.Print(err)
		return 0, err
	}

	return count, nil
}

func (s *DBUserService) Exists(where string, args ...interface{}) (bool, error) {
	count, err := s.Count(where, args...)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *DBUserService) Get(where string, args ...interface{}) (*User, error) {
	u := User{}
	q := "SELECT * FROM users WHERE " + where

	if err := s.db.Get(&u, q, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}

func (s *DBUserService) List(query string, args ...interface{}) ([]User, error) {
	q := "SELECT * FROM users " + query

	rows, err := s.db.Queryx(q, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := []User{}

	for rows.Next() {
		u := User{}
		rows.StructScan(&u)
		results = append(results, u)
	}

	return results, nil
}
