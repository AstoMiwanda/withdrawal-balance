package model

import "time"

type User struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name" validate:"required"`
	Phone     string     `json:"phone" db:"phone" validate:"required"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
