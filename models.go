package api

import "time"

type BaseModel struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// swagger:model
type User struct {
	BaseModel

	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
}
