package model

import (
	"time"
)

type User struct {
	ID           string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	TotalMatches int       `json:"matches"`
	TotalMmr     string    `json:"mmr"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
