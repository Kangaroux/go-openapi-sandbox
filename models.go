package api

import "time"

type BaseModel struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// swagger:model
type User struct {
	BaseModel

	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"-"`
	PasswordSalt string `json:"-"`
}
