package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	TotalMatches int       `json:"matches"`
	TotalMmr     string    `json:"mmr"`
	CreatedAt    time.Time `json:"created_at"`
}
